package interest

import (
	"github.com/metabloxStaking/dao"
	"math/big"
	"strconv"

	"github.com/metabloxStaking/models"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func TestTruncateToHour(t *testing.T) {
	time1, _ := time.Parse("2006-01-02 15:04:05", "2022-05-27 13:31:00")
	expected1, _ := time.Parse("2006-01-02 15:04:05", "2022-05-27 13:00:00")
	assert.Equal(t, expected1, time1)
}

func TestCalculateCurrentAPY(t *testing.T) {
	tests := []struct {
		name           string
		topUpLimit     *big.Int
		lockUpPeriod   int
		defaultAPY     float64
		totalPrincipal *big.Int
		expected       string
	}{
		{
			name:           "total principal 400",
			topUpLimit:     big.NewInt(500000),
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: big.NewInt(400),
			expected:       "250",
		},
		{
			name:           "total principal 900",
			topUpLimit:     big.NewInt(500000),
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: big.NewInt(900),
			expected:       "111.1111111",
		},
		{
			name:           "total principal 1500",
			topUpLimit:     big.NewInt(500000),
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: big.NewInt(1500),
			expected:       "66.66666667",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product := &models.StakingProduct{
				TopUpLimit:   tt.topUpLimit,
				LockUpPeriod: tt.lockUpPeriod,
				DefaultAPY:   tt.defaultAPY,
			}
			assert.Equal(t, tt.expected, CalculateCurrentAPY(product, tt.totalPrincipal).String())
		})
	}
}

func TestGetOrderInterestList(t *testing.T) {
	err := dao.InitTestDB()
	assert.Nil(t, err)
	defer dao.CleanupTestDB()

	order := &models.Order{
		ProductID:    "1",
		UserDID:      "test",
		Type:         "Pending",
		Amount:       big.NewInt(400000000),
		StringAmount: "400000000",
	}
	id, err := dao.CreateOrder(order)
	assert.Nil(t, err)

	err = dao.InsertPrincipalUpdate(order.ProductID, order.StringAmount)
	assert.Nil(t, err)

	txInfo := &models.TXInfo{
		OrderID:         strconv.Itoa(id),
		TXCurrencyType:  "",
		TXType:          "BuyIn",
		TXHash:          nil,
		Principal:       order.Amount,
		StringPrincipal: order.Amount.String(),
		Interest:        big.NewInt(0),
		StringInterest:  "0",
		UserAddress:     "",
		RedeemableTime:  "2022-01-01 00:00:00",
	}
	err = dao.SubmitBuyin(txInfo)
	assert.Nil(t, err)

	until := time.Now().UTC().Add(time.Hour)
	result, err := GetOrderInterestList(strconv.Itoa(id), until)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "11574074", result[0].InterestGain.String())
}
