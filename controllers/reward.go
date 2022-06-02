package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/foundationdao"
)

func GetRewardHistory(c *gin.Context) (float64, error) {
	userDID := c.Param("did")

	exchangeList, err := foundationdao.GetSeedHistory(userDID)
	if err != nil {
		return 0, err
	}

	redeemedToken := 0.0

	for _, exchange := range exchangeList {
		redeemedToken += exchange.Amount
	}

	return redeemedToken, nil
}
