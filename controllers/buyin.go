package controllers

import (
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

	txInfo := models.NewTXInfo(input.OrderID, models.CurrencyTypeMBLX, models.TXTypeBuyin, input.TxHash, order.Amount, 0, order.UserAddress, time.Now().AddDate(0, 0, 179).Truncate(24*time.Hour).Format("2006-01-02 15:04:05.000"))

	err = dao.SubmitBuyin(txInfo)
	if err != nil {
		return nil, err
	}

	date, err := dao.GetTXCreateDate(input.TxHash)
	if err != nil {
		return nil, err
	}

	output := models.NewSubmitBuyinOutput(product.ProductName, order.Amount, date, txInfo.TXCurrencyType, txInfo.UserAddress)
	return output, err
}
