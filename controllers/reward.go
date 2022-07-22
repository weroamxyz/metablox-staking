package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/errval"
	"github.com/shopspring/decimal"
)

//get all seed exchanges where the user was the specified did, and return the total amount of MBLX redeemed
func GetRewardHistory(c *gin.Context) (string, error) {
	userDID := c.Param("did")

	exchangeList, err := dao.GetSeedHistory(userDID)
	if err != nil {
		return "", err
	}

	redeemedToken := decimal.NewFromInt(0)

	for _, exchange := range exchangeList {
		amount, err := decimal.NewFromString(exchange.Amount)
		if err != nil {
			return "", errval.ErrAmountNotNumber
		}
		redeemedToken = redeemedToken.Add(amount)
	}
	convertedRedeemedAmount := redeemedToken.Div(decimal.NewFromInt(1000000))
	return convertedRedeemedAmount.String(), nil
}
