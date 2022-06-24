package controllers

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"time"

	logger "github.com/sirupsen/logrus"

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

	totalPrincipal := decimal.NewFromInt(0)
	principalUpdate, err := dao.GetLatestPrincipalUpdate(input.ProductID)
	if err == nil {
		totalPrincipal = principalUpdate.TotalPrincipal
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	product, err := dao.GetProductInfoByID(input.ProductID)
	if err != nil {
		return nil, err
	}

	floatAmount, err := strconv.ParseFloat(input.Amount, 64)
	if err != nil {
		return nil, err
	}
	bigAmount := models.MBLXToMinimumUnit(floatAmount)

	if bigAmount.Cmp(product.MinOrderValue) == -1 {
		return nil, errval.ErrOrderAmountTooLow
	}
	if totalPrincipal.Add(bigAmount).Cmp(product.TopUpLimit) == 1 {
		return nil, errval.ErrOrderExceedsTopUpLimit
	}

	paymentAddress, err := dao.GetPaymentAddress(input.ProductID)
	if err != nil {
		return nil, err
	}

	newOrder := models.NewOrder(input.ProductID, input.UserDID, models.OrderTypePending, paymentAddress, bigAmount, input.UserAddress)

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

	currentInterest := interestInfo.AccumulatedInterest.Sub(interestInfo.TotalInterestGained)

	amount := currentInterest.Add(order.Amount)
	tx, err := contract.RedeemOrder(order.UserAddress, amount)
	if err != nil {
		return nil, err
	}
	txData, _ := tx.MarshalJSON()
	logger.Infof("tx %s send,detail:%s", tx.Hash().Hex(), string(txData))

	txInfo := models.NewTXInfo(orderID, models.CurrencyTypeMBLX, models.TxTypeOrderClosure, tx.Hash().Hex(), order.Amount, currentInterest, userAddress, redeemableDate)

	err = dao.RedeemOrder(txInfo, interestInfo.AccumulatedInterest.String())
	if err != nil {
		return nil, err
	}

	convertedAmount := models.MinimumUnitToMBLX(amount)
	time := strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output := models.NewRedeemOrderOutput(productName, strconv.FormatFloat(convertedAmount, 'f', -1, 64), time, userAddress, models.CurrencyTypeMBLX, tx.Hash().Hex())

	// record change in staking pool's total principal
	newPrincipal := models.NewPrincipalUpdate()
	oldPrincipal, err := dao.GetLatestPrincipalUpdate(order.ProductID)
	if err != nil {
		return nil, err
	}
	newPrincipal.TotalPrincipal = oldPrincipal.TotalPrincipal.Sub(order.Amount)

	err = dao.InsertPrincipalUpdate(order.ProductID, newPrincipal.TotalPrincipal.String())
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

	currentInterest := interestInfo.AccumulatedInterest.Sub(interestInfo.TotalInterestGained)

	valid, err := dao.CompareMinimumInterest(orderID, currentInterest.String())
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
	txInfo := models.NewTXInfo(orderID, models.CurrencyTypeMBLX, models.TxTypeInterestOnly, tx.Hash().Hex(), decimal.NewFromInt(0), currentInterest, userAddress, time.Now().Format("2006-01-02 15:04:05.000"))

	productName, err := dao.GetProductNameForOrder(orderID)
	if err != nil {
		return nil, err
	}

	convertedInterest := models.MinimumUnitToMBLX(currentInterest)
	time := strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output := models.NewRedeemOrderOutput(productName, strconv.FormatFloat(convertedInterest, 'f', -1, 64), time, userAddress, models.CurrencyTypeMBLX, tx.Hash().Hex())

	err = dao.RedeemInterest(txInfo)
	if err != nil {
		return nil, err
	}

	return output, nil
}
