package controllers

import (
	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/errval"
)

//get all seed exchanges where the user was the specified did, and return the total amount of MBLX redeemed
func GetRewardHistory(c *gin.Context) (string, error) {
	userDID := c.Param("did")

	exchangeList, err := dao.GetSeedHistory(userDID)
	if err != nil {
		return "", err
	}

	redeemedToken := big.NewInt(0)

	for _, exchange := range exchangeList {
		bigAmount, success := big.NewInt(0).SetString(exchange.Amount, 10)
		if !success {
			return "", errval.ErrAmountNotNumber
		}
		redeemedToken.Add(redeemedToken, bigAmount)
	}

	convertedRedeemedAmount := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(redeemedToken), big.NewFloat(1000000))
	return convertedRedeemedAmount.String(), nil
}
