package interest

import (
	"math/big"
	"strconv"

	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

const floatErrorTolerance = 1e-6

func TestTruncateToHour(t *testing.T) {
	time1, _ := time.Parse("2006-01-02 15:04:05", "2022-05-27 13:31:00")
	expected1, _ := time.Parse("2006-01-02 15:04:05", "2022-05-27 13:00:00")
	assert.Equal(t, expected1, TruncateToHour(time1))
}

func TestCalculateCurrentAPY(t *testing.T) {
	tests := []struct {
		name           string
		topUpLimit     *big.Int
		lockUpPeriod   int
		defaultAPY     float64
		totalPrincipal *big.Int
		expected       float64
	}{
		{
			name:           "total principal 400",
			topUpLimit:     big.NewInt(500000),
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: big.NewInt(400),
			expected:       250.00,
		},
		{
			name:           "total principal 900",
			topUpLimit:     big.NewInt(500000),
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: big.NewInt(900),
			expected:       111.111111,
		},
		{
			name:           "total principal 1500",
			topUpLimit:     big.NewInt(500000),
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: big.NewInt(1500),
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

func TestUpdateOrderInterest(t *testing.T) {
	err := dao.InitTestDB()
	assert.Nil(t, err)
	defer dao.CleanupTestDB()

	product, err := dao.GetProductInfoByID("1")
	assert.Nil(t, err)

	order := &models.Order{
		ProductID: product.ID,
		UserDID:   "test",
		Type:      "Pending",
		Amount:    big.NewInt(400),
	}
	id, err := dao.CreateOrder(order)
	assert.Nil(t, err)
	orderID := strconv.Itoa(id)

	now := time.Now().UTC()
	err = dao.InsertPrincipalUpdate(product.ID, order.Amount.String())
	assert.Nil(t, err)

	principalUpdates, err := dao.GetPrincipalUpdates(product.ID)
	assert.Nil(t, err)

	txInfo := &models.TXInfo{
		OrderID:        orderID,
		TXCurrencyType: "",
		TXType:         "BuyIn",
		TXHash:         nil,
		Principal:      order.Amount,
		Interest:       big.NewInt(0),
		UserAddress:    "",
		RedeemableTime: "2022-01-01 00:00:00",
	}
	err = dao.SubmitBuyin(txInfo)
	assert.Nil(t, err)

	until := now.Add(time.Hour)

	err = UpdateOrderInterest(orderID, product, principalUpdates, until)
	assert.Nil(t, err)
	result, err := dao.GetSortedOrderInterestListUntilDate(orderID, until.Format(TimeFormat))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.InEpsilon(t, 11.574074, result[0].InterestGain, floatErrorTolerance)

	until = now.Add(time.Hour * 10)

	err = UpdateOrderInterest(orderID, product, principalUpdates, until)
	assert.Nil(t, err)
	result, err = dao.GetSortedOrderInterestListUntilDate(orderID, until.Format(TimeFormat))
	assert.Nil(t, err)
	assert.Equal(t, 10, len(result))
	assert.InEpsilon(t, 11.574074, result[9].InterestGain, floatErrorTolerance)
	assert.InEpsilon(t, 115.740741, result[9].TotalInterestGain, floatErrorTolerance)
}

func TestOrderInterest_MultipleUsers(t *testing.T) {
	events := []struct {
		orderID  string
		amount   *big.Int
		hour     int  // number of hours since product start
		isRedeem bool // false = buy-in, true = redeem
	}{
		{"1", big.NewInt(200000), 271, false},
		{"2", big.NewInt(125000), 1525, false},
		{"3", big.NewInt(125000), 2681, false},
		{"4", big.NewInt(50000), 3634, false},
		{"1", big.NewInt(0), 4591, true},
		{"2", big.NewInt(0), 5845, true},
		{"3", big.NewInt(0), 7001, true},
	}

	err := dao.InitTestDB()
	assert.Nil(t, err)
	defer dao.CleanupTestDB()

	product, err := dao.GetProductInfoByID("1")
	assert.Nil(t, err)
	productStart, err := time.Parse(TimeFormat, product.StartDate)
	assert.Nil(t, err)

	prevUpdate := productStart
	for _, event := range events {
		currTime := productStart.Add(time.Hour * time.Duration(event.hour))
		if currTime.After(prevUpdate) {
			updateAllOrderInterest(currTime) // update before each new order buyin/redeem
			prevUpdate = currTime
		}
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

		{orderID: "1", hour: 1526, expectedLen: 1255, expectedInterestGain: 7.122507123},
		{orderID: "2", hour: 1526, expectedLen: 1, expectedInterestGain: 4.451566952},

		{orderID: "1", hour: 2682, expectedLen: 2411, expectedInterestGain: 5.144032922},
		{orderID: "2", hour: 2682, expectedLen: 1157, expectedInterestGain: 3.215020576},
		{orderID: "3", hour: 2682, expectedLen: 1, expectedInterestGain: 3.215020576},

		{orderID: "1", hour: 3635, expectedLen: 3364, expectedInterestGain: 4.62962963},
		{orderID: "2", hour: 3635, expectedLen: 2110, expectedInterestGain: 2.893518519},
		{orderID: "3", hour: 3635, expectedLen: 954, expectedInterestGain: 2.893518519},
		{orderID: "4", hour: 3635, expectedLen: 1, expectedInterestGain: 1.157407407},

		// order 1 redeemed, should stay at length 4320 from now on
		{orderID: "1", hour: 4592, expectedLen: 4320, expectedInterestGain: 4.62962963},
		{orderID: "2", hour: 4592, expectedLen: 3067, expectedInterestGain: 4.822530864},
		{orderID: "3", hour: 4592, expectedLen: 1911, expectedInterestGain: 4.822530864},
		{orderID: "4", hour: 4592, expectedLen: 958, expectedInterestGain: 1.929012346},

		// all redeemed orders should stay at length 4320
		{orderID: "1", hour: 5846, expectedLen: 4320, expectedInterestGain: 4.62962963},
		{orderID: "2", hour: 5846, expectedLen: 4320, expectedInterestGain: 4.822530864},
		{orderID: "3", hour: 5846, expectedLen: 3165, expectedInterestGain: 8.267195767},
		{orderID: "4", hour: 5846, expectedLen: 2212, expectedInterestGain: 3.306878307},

		{orderID: "1", hour: 7002, expectedLen: 4320, expectedInterestGain: 4.62962963},
		{orderID: "2", hour: 7002, expectedLen: 4320, expectedInterestGain: 4.822530864},
		{orderID: "3", hour: 7002, expectedLen: 4320, expectedInterestGain: 8.267195767},
		{orderID: "4", hour: 7002, expectedLen: 3368, expectedInterestGain: 11.57407407},

		// order 4 was never redeemed, so it should continue gaining interest in the next product term
		{orderID: "1", hour: 7002, expectedLen: 4320, expectedInterestGain: 4.62962963},
		{orderID: "2", hour: 7002, expectedLen: 4320, expectedInterestGain: 4.822530864},
		{orderID: "3", hour: 7002, expectedLen: 4320, expectedInterestGain: 8.267195767},
		{orderID: "4", hour: 12274, expectedLen: 8640, expectedInterestGain: 11.57407407},
	}

	// not using tt.Run because we don't want to set up the purchases multiple times
	for _, tt := range tests {
		currTime := productStart.Add(time.Hour * time.Duration(tt.hour))
		if currTime.After(prevUpdate) {
			updateAllOrderInterest(currTime) // update before each new order buyin/redeem
			prevUpdate = currTime
		}
		result, err := dao.GetSortedOrderInterestListUntilDate(tt.orderID, currTime.Format(TimeFormat))
		assert.Nil(t, err)
		if assert.Equal(t, tt.expectedLen, len(result)) {
			if !assert.InEpsilon(t, tt.expectedInterestGain, result[len(result)-1].InterestGain, floatErrorTolerance) {
				logger.Warnf("test order %s at hour %d failed. Expected interest = %f, actual interest = %f", tt.orderID, tt.hour, tt.expectedInterestGain, result[len(result)-1].InterestGain)
			}
		} else {
			logger.Warnf("test order %s at hour %d failed", tt.orderID, tt.hour)
		}
	}
}
