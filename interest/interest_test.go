package interest

import (
	"math/big"

	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
	"github.com/stretchr/testify/assert"

	"strconv"
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
		topUpLimit     float64
		lockUpPeriod   int
		defaultAPY     float64
		totalPrincipal float64
		expected       string
	}{
		{
			name:           "total principal 400",
			topUpLimit:     500000,
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: 400,
			expected:       "250.00",
		},
		{
			name:           "total principal 900",
			topUpLimit:     500000,
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: 900,
			expected:       "111.111111",
		},
		{
			name:           "total principal 1500",
			topUpLimit:     500000,
			lockUpPeriod:   180,
			defaultAPY:     0.2,
			totalPrincipal: 1500,
			expected:       "66.666667",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product := &models.StakingProduct{
				TopUpLimit:   tt.topUpLimit,
				LockUpPeriod: tt.lockUpPeriod,
				DefaultAPY:   tt.defaultAPY,
			}
			assert.Equal(t, tt.expected, FormatFloat(CalculateCurrentAPY(product, big.NewInt(int64(tt.totalPrincipal)))))
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
		StringAmount: "400",
	}
	id, err := dao.CreateOrder(order)
	assert.Nil(t, err)

	err = dao.InsertPrincipalUpdate(order.ProductID, order.Amount.String())
	assert.Nil(t, err)

	txInfo := &models.TXInfo{
		OrderID:        strconv.Itoa(id),
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

	until := time.Now().Add(time.Hour)
	result, err := GetOrderInterestList(strconv.Itoa(id), until)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "11.574074", result[0].InterestGain.String())
}
