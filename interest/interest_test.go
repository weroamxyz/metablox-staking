package interest

import (
	"fmt"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"strconv"
	"testing"
	"time"
)

const floatErrorTolerance = 1e-6

func TestTruncateToHour(t *testing.T) {
	time1, _ := time.Parse("2006-01-02 15:04:05", "2022-05-27 13:31:00")
	expected1, _ := time.Parse("2006-01-02 15:04:05", "2022-05-27 13:00:00")
	assert.Equal(t, expected1, time1)
}

func TestCalculateCurrentAPY(t *testing.T) {
	tests := []struct {
		name           string
		topUpLimit     float64
		lockUpPeriod   int
		defaultAPY     float64
		totalPrincipal float64
		expected       float64
	}{
		{
			name:           "total principal 400",
			topUpLimit:     500000,
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: 400,
			expected:       250.00,
		},
		{
			name:           "total principal 900",
			topUpLimit:     500000,
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: 900,
			expected:       111.111111,
		},
		{
			name:           "total principal 1500",
			topUpLimit:     500000,
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: 1500,
			expected:       66.666667,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product := &models.StakingProduct{
				TopUpLimit:   tt.topUpLimit,
				LockUpPeriod: tt.lockUpPeriod,
				DefaultAPY:   tt.defaultAPY,
			}
			assert.InEpsilon(t, tt.expected, CalculateCurrentAPY(product, tt.totalPrincipal), floatErrorTolerance)
		})
	}
}

func TestGetOrderInterestList(t *testing.T) {
	err := dao.InitTestDB()
	assert.Nil(t, err)
	defer dao.CleanupTestDB()

	order := &models.Order{
		ProductID: "1",
		UserDID:   "test",
		Type:      "Pending",
		Amount:    400,
	}
	id, err := dao.CreateOrder(order)
	assert.Nil(t, err)

	err = dao.InsertPrincipalUpdate(order.ProductID, order.Amount)
	assert.Nil(t, err)

	txInfo := &models.TXInfo{
		OrderID:        strconv.Itoa(id),
		TXCurrencyType: "",
		TXType:         "BuyIn",
		TXHash:         nil,
		Principal:      order.Amount,
		Interest:       0,
		UserAddress:    "",
		RedeemableTime: "2022-01-01 00:00:00",
	}
	err = dao.SubmitBuyin(txInfo)
	assert.Nil(t, err)

	until := time.Now().Add(time.Hour)
	result, err := GetOrderInterestList(strconv.Itoa(id), until)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.InEpsilon(t, 11.574074, result[0].InterestGain, floatErrorTolerance)

	until = until.Add(time.Hour * 9)
	result, err = GetOrderInterestList(strconv.Itoa(id), until)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(result))
	assert.InEpsilon(t, 11.574074, result[9].InterestGain, floatErrorTolerance)
	assert.InEpsilon(t, 115.740741, result[9].TotalInterestGain, floatErrorTolerance)
}

func TestOrderInterest_MultipleUsers(t *testing.T) {
	events := []struct {
		orderID  string
		amount   float64
		hour     int  // number of hours since product start
		isRedeem bool // false = buy-in, true = redeem
	}{
		{"1", 200000, 271, false},
		{"2", 100000, 800, false},
		{"3", 100000, 1525, false},
		{"4", 50000, 2433, false},
		{"5", 20000, 2681, false},
		{"6", 20000, 3347, false},
		{"7", 10000, 3634, false},
		{"1", 0, 4591, true},
		{"2", 0, 5120, true},
		{"3", 0, 5845, true},
		{"4", 0, 6753, true},
		{"5", 0, 7001, true},
		{"6", 0, 7667, true},
	}

	err := dao.InitTestDB()
	assert.Nil(t, err)
	defer dao.CleanupTestDB()

	product, err := dao.GetProductInfoByID("1")
	assert.Nil(t, err)
	productStart, err := time.Parse(TimeFormat, product.StartDate)
	assert.Nil(t, err)

	for _, event := range events {
		currTime := productStart.Add(time.Hour * time.Duration(event.hour))
		if event.hour == 4591 {
			fmt.Println("hi")
		}
		updateAllOrderInterest(currTime) // update before each new order buyin/redeem
		assert.Nil(t, err)
		if !event.isRedeem { // order buy-in
			order := &models.Order{
				ProductID: "1",
				UserDID:   "test",
				Type:      "Pending",
				Amount:    event.amount,
			}
			id, err := dao.BuyinTestOrderWithDate(order, currTime.Format(TimeFormat))
			assert.Nil(t, err)
			assert.Equal(t, event.orderID, id)
		} else {
			err := dao.RedeemTestOrderWithDate(event.orderID, currTime.Format(TimeFormat))
			assert.Nil(t, err)
		}
	}

	// Each test checks the number of entries and the interest gain for all active orders one hour after each new order or redemption.
	tests := []struct {
		orderID              string
		hour                 int
		expectedLen          int
		expectedInterestGain float64
	}{
		{orderID: "1", hour: 272, expectedLen: 1, expectedInterestGain: 11.57407407},

		{orderID: "1", hour: 801, expectedLen: 530, expectedInterestGain: 7.716049383},
		{orderID: "2", hour: 801, expectedLen: 1, expectedInterestGain: 3.858024691},

		{orderID: "1", hour: 1526, expectedLen: 1255, expectedInterestGain: 5.787037037},
		{orderID: "2", hour: 1526, expectedLen: 726, expectedInterestGain: 2.893518519},
		{orderID: "3", hour: 1526, expectedLen: 1, expectedInterestGain: 2.893518519},

		{orderID: "1", hour: 2434, expectedLen: 2163, expectedInterestGain: 5.144032922},
		{orderID: "2", hour: 2434, expectedLen: 1634, expectedInterestGain: 2.572016461},
		{orderID: "3", hour: 2434, expectedLen: 909, expectedInterestGain: 2.572016461},
		{orderID: "4", hour: 2434, expectedLen: 1, expectedInterestGain: 1.28600823},

		{orderID: "1", hour: 2682, expectedLen: 2411, expectedInterestGain: 4.925137904},
		{orderID: "2", hour: 2682, expectedLen: 1882, expectedInterestGain: 2.462568952},
		{orderID: "3", hour: 2682, expectedLen: 1157, expectedInterestGain: 2.462568952},
		{orderID: "4", hour: 2682, expectedLen: 249, expectedInterestGain: 1.231284476},
		{orderID: "5", hour: 2682, expectedLen: 1, expectedInterestGain: 0.4925137904},

		{orderID: "1", hour: 3348, expectedLen: 3077, expectedInterestGain: 4.724111867},
		{orderID: "2", hour: 3348, expectedLen: 2548, expectedInterestGain: 2.362055933},
		{orderID: "3", hour: 3348, expectedLen: 1823, expectedInterestGain: 2.362055933},
		{orderID: "4", hour: 3348, expectedLen: 915, expectedInterestGain: 1.181027967},
		{orderID: "5", hour: 3348, expectedLen: 667, expectedInterestGain: 0.4724111867},
		{orderID: "6", hour: 3348, expectedLen: 1, expectedInterestGain: 0.4724111867},

		{orderID: "1", hour: 3635, expectedLen: 3364, expectedInterestGain: 4.62962963},
		{orderID: "2", hour: 3635, expectedLen: 2835, expectedInterestGain: 2.314814815},
		{orderID: "3", hour: 3635, expectedLen: 2110, expectedInterestGain: 2.314814815},
		{orderID: "4", hour: 3635, expectedLen: 1202, expectedInterestGain: 1.157407407},
		{orderID: "5", hour: 3635, expectedLen: 954, expectedInterestGain: 0.462962963},
		{orderID: "6", hour: 3635, expectedLen: 288, expectedInterestGain: 0.462962963},
		{orderID: "7", hour: 3635, expectedLen: 1, expectedInterestGain: 0.2314814815},

		// order 1 redeemed, should stay at length 4320 from now on
		{orderID: "1", hour: 4592, expectedLen: 4320, expectedInterestGain: 4.62962963},
		{orderID: "2", hour: 4592, expectedLen: 3792, expectedInterestGain: 3.858024691},
		{orderID: "3", hour: 4592, expectedLen: 3067, expectedInterestGain: 3.858024691},
		{orderID: "4", hour: 4592, expectedLen: 2159, expectedInterestGain: 1.929012346},
		{orderID: "5", hour: 4592, expectedLen: 1911, expectedInterestGain: 0.7716049383},
		{orderID: "6", hour: 4592, expectedLen: 1245, expectedInterestGain: 0.7716049383},
		{orderID: "7", hour: 4592, expectedLen: 958, expectedInterestGain: 0.3858024691},
	}

	// not using tt.Run because we don't want to set up the purchases multiple times
	for _, tt := range tests {
		logger.Infof("checking order %s at hour %d", tt.orderID, tt.hour)
		currTimeStr := productStart.Add(time.Hour * time.Duration(tt.hour)).Format(TimeFormat)
		result, err := dao.GetSortedOrderInterestListUntilDate(tt.orderID, currTimeStr)
		assert.Nil(t, err)
		assert.Equal(t, tt.expectedLen, len(result))
		if !assert.InEpsilon(t, tt.expectedInterestGain, result[len(result)-1].InterestGain, floatErrorTolerance) {
			logger.Infof("expected = %f, actual = %f", tt.expectedInterestGain, result[len(result)-1].InterestGain)
		}
	}
}
