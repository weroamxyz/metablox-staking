package interest

import (
	"fmt"
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
			expected:       "250.000000",
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
			assert.Equal(t, tt.expected, fmt.Sprintf("%.6f", CalculateCurrentAPY(product, tt.totalPrincipal)))
		})
	}
}

func TestGetOrderInterestList(t *testing.T) {

}

func TestGetAllOrderInterest(t *testing.T) {

}
