package controllers

import (
	logger "github.com/sirupsen/logrus"
	"math"
	"strconv"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/contract"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
)

func CreateOrder(c *gin.Context) (*models.OrderOutput, error) {
	input := models.CreateOrderInput()
	err := c.BindJSON(input)
	if err != nil {
		return nil, err
	}

	valid := did.IsDIDValid(did.SplitDIDString(input.UserDID))
	if !valid {
		return nil, errval.ErrBadDID
	}

	newOrder := models.NewOrder(input.ProductID, input.UserDID, models.OrderTypePending, "placeholder", input.Amount, input.UserAddress)

	orderID, err := dao.CreateOrder(newOrder)
	if err != nil {
		return nil, err
	}

	output := models.NewOrderOutput(strconv.Itoa(orderID), newOrder.PaymentAddress)

	return output, nil
}

func RedeemOrder(c *gin.Context) (*models.RedeemOrderOuput, error) {
	orderID := c.Param("id")

	redeemableDate, err := dao.GetOrderRedeemableDate(orderID)
	if err != nil {
		return nil, err
	}
	redeemableTime, err := time.Parse("2006-01-02 15:04:05", redeemableDate)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return nil, err
	}

	elapsedTime := time.Since(redeemableTime)
	elapsedDays := int(math.Floor(elapsedTime.Hours() / 24))
	if elapsedDays != 0 {
		return nil, errval.ErrEarlyOrderRedeem
	}

	userAddress, err := dao.GetUserAddressByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	interestInfo, err := dao.GetInterestInfoByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	productName, err := dao.GetProductNameForOrder(orderID)
	if err != nil {
		return nil, err
	}

	order, err := dao.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	amount := (interestInfo.AccumulatedInterest - interestInfo.TotalInterestGained) + order.Amount
	tx, err := contract.RedeemOrder(order.UserAddress, amount)
	if err != nil {
		return nil, err
	}
	txData, _ := tx.MarshalJSON()
	logger.Infof("tx %s send,detail:%s", tx.Hash().Hex(), string(txData))

	txInfo := models.NewTXInfo(orderID, models.CurrencyTypeMBLX, models.TxTypeOrderClosure, tx.Hash().Hex(), 0, 0, userAddress, redeemableDate)

	dao.RedeemOrder(txInfo, interestInfo.AccumulatedInterest)
	err = dao.UploadTransaction(txInfo)
	if err != nil {
		return nil, err
	}

	time := strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output := models.NewRedeemOrderOutput(productName, amount, time, userAddress, models.CurrencyTypeMBLX, tx.Hash().Hex())

	// record change in staking pool's total principal
	newPrincipal := models.NewPrincipalUpdate()
	oldPrincipal, err := dao.GetLatestPrincipalUpdate(order.ProductID)
	if err != nil {
		return nil, err
	}
	newPrincipal.TotalPrincipal = oldPrincipal.TotalPrincipal - order.Amount

	err = dao.InsertPrincipalUpdate(order.ProductID, newPrincipal.TotalPrincipal)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func RedeemInterest(c *gin.Context) (*models.RedeemOrderOuput, error) {
	orderID := c.Param("id")

	interestInfo, err := dao.GetInterestInfoByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	currentInterest := interestInfo.AccumulatedInterest - interestInfo.TotalInterestGained

	valid, err := dao.CompareMinimumInterest(orderID, currentInterest)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errval.ErrNotEnoughInterest
	}

	userAddress, err := dao.GetUserAddressByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	tx, err := contract.RedeemOrder(userAddress, currentInterest)
	if err != nil {
		return nil, err
	}
	txData, _ := tx.MarshalJSON()
	logger.Infof("tx %s send,detail:%s"+tx.Hash().Hex(), string(txData))
	txInfo := models.NewTXInfo(orderID, models.CurrencyTypeMBLX, models.TxTypeInterestOnly, tx.Hash().Hex(), 0, 0, userAddress, time.Now().Format("2006-01-02 15:04:05.000"))

	productName, err := dao.GetProductNameForOrder(orderID)
	if err != nil {
		return nil, err
	}

	time := strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output := models.NewRedeemOrderOutput(productName, currentInterest, time, userAddress, models.CurrencyTypeMBLX, tx.Hash().Hex())

	err = dao.RedeemOrder(txInfo, interestInfo.AccumulatedInterest)
	if err != nil {
		return nil, err
	}

	return output, nil
}
