package controllers

import (
	"database/sql"
	"math/big"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/contract"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
)

func SubmitBuyin(c *gin.Context) (*models.SubmitBuyinOutput, error) {
	input := models.CreateSubmitBuyinInput()
	err := c.BindJSON(input)
	if err != nil {
		return nil, err
	}

	exists, err := dao.CheckIfTXExists(input.TxHash)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errval.ErrExistingTXHash
	}

	order, err := dao.GetOrderByID(input.OrderID)
	if err != nil {
		return nil, err
	}

	err = contract.CheckIfTransactionMatchesOrder(input.TxHash, order)
	if err != nil {
		return nil, err
	}

	product, err := dao.GetProductInfoByID(order.ProductID)
	if err != nil {
		return nil, err
	}

	txInfo := models.NewTXInfo(input.OrderID, models.CurrencyTypeMBLX, models.TxTypeBuyIn, input.TxHash, order.Amount, big.NewInt(0), order.UserAddress, time.Now().AddDate(0, 0, 179).Truncate(24*time.Hour).Format("2006-01-02 15:04:05.000"))

	err = dao.SubmitBuyin(txInfo)
	if err != nil {
		return nil, err
	}

	date, err := dao.GetTXCreateDate(input.TxHash)
	if err != nil {
		return nil, err
	}

	order.SetMBLXValues()
	output := models.NewSubmitBuyinOutput(product.ProductName, strconv.FormatFloat(order.MBLXAmount, 'f', -1, 64), date, txInfo.UserAddress, txInfo.TXCurrencyType)

	// record change in staking pool's total principal
	newPrincipal := models.NewPrincipalUpdate()
	oldPrincipal, err := dao.GetLatestPrincipalUpdate(product.ID)
	if err == nil {
		newPrincipal.TotalPrincipal = big.NewInt(0).Add(oldPrincipal.TotalPrincipal, txInfo.Principal)
	} else if err == sql.ErrNoRows {
		newPrincipal.TotalPrincipal = txInfo.Principal
	} else {
		return nil, err
	}

	err = dao.InsertPrincipalUpdate(product.ID, newPrincipal.TotalPrincipal.String())
	if err != nil {
		return nil, err
	}

	return output, err
}
