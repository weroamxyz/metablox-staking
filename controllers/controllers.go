package controllers

import (
	"time"

	"github.com/metabloxStaking/interest"

	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/contract"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/errval"
	logger "github.com/sirupsen/logrus"
)

//check if DID is formatted properly and is in DID registry
func ValidateDID(userDID string) error {
	splitDID := did.SplitDIDString(userDID)
	valid := did.IsDIDValid(splitDID)
	if !valid {
		return errval.ErrBadDID
	}
	err := contract.CheckForRegisteredDID(splitDID[2])
	if err != nil {
		return err
	}
	return nil
}

//get staking product with specified id. Also calculate its current APY
func GetProductInfoByIDHandler(c *gin.Context) {
	productID := c.Param("id")
	product, err := dao.GetProductInfoByID(productID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	principalUpdate, err := dao.GetLatestPrincipalUpdate(product.ID)
	if err != nil {
		product.CurrentAPY = product.DefaultAPY
	} else {
		product.CurrentAPY, _ = interest.CalculateCurrentAPY(product, principalUpdate.TotalPrincipal).Float64()
	}
	ResponseSuccess(c, product)
}

//get all staking products. Also calculate their current APYs
func GetAllProductInfoHandler(c *gin.Context) {
	products, err := dao.GetAllProductInfo()
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	for _, product := range products {
		principalUpdate, err := dao.GetLatestPrincipalUpdate(product.ID)
		if err != nil {
			product.CurrentAPY = product.DefaultAPY
		} else {
			product.CurrentAPY, _ = interest.CalculateCurrentAPY(product, principalUpdate.TotalPrincipal).Float64()
		}
	}
	ResponseSuccess(c, products)
}

func CreateOrderHandler(c *gin.Context) {
	output, err := CreateOrder(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
}

func SubmitBuyinHandler(c *gin.Context) {
	output, err := SubmitBuyin(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
}

func GetStakingRecordsHandler(c *gin.Context) {
	records, err := GetStakingRecords(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, records)
}

//get all transactions associated with the specified order id
func GetTransactionsByOrderIDHandler(c *gin.Context) {
	orderID := c.Param("id")
	transactions, err := dao.GetTransactionsByOrderID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, transactions)
}

//get all transactions associated with the specified user did
func GetTransactionsByUserDIDHandler(c *gin.Context) {
	userDID := c.Param("did")

	transactions, err := dao.GetTransactionsByUserDID(userDID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, transactions)
}

//get all order interest entries that have happened up until now, sorted in ascending order of recency
func GetOrderInterestHandler(c *gin.Context) {
	orderID := c.Param("id")
	transactions, err := dao.GetSortedOrderInterestListUntilDate(orderID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, transactions)
}

func RedeemOrderHandler(c *gin.Context) {
	output, err := RedeemOrder(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
}

func RedeemInterestHandler(c *gin.Context) {
	output, err := RedeemInterest(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
}

func GetMinerListHandler(c *gin.Context) {
	minerList, err := GetMinerList(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, minerList)
}

func GetMinerByIDHandler(c *gin.Context) {
	miner, err := GetMinerByID(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, miner)
}

func GetExchangeRateHandler(c *gin.Context) {
	exchangeRate, err := GetExchangeRate(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, exchangeRate)
}

func GetRewardHistoryHandler(c *gin.Context) {
	redeemedToken, err := GetRewardHistory(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, redeemedToken)
}

func GetNonceHandler(c *gin.Context) {
	output, err := GetNonce(c)
	if err != nil {
		logger.Error(err)
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
}

func ActivateExchangeHandler(c *gin.Context) {
	err := ActivateExchange(c)
	if err != nil {
		logger.Error(err)
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, "")
}

func NewSeedExchangeHandler(c *gin.Context) {
	output, err := NewExchangeSeed(c)
	if err != nil {
		logger.Error(err)
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
}

func GetWifiAccessInfoHandler(c *gin.Context) {
	wifi, err := GetWifiByAccount(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, wifi)
}

func GetMinerByBSSIDHandler(c *gin.Context) {
	miner, err := GetMinerByBSSID(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, miner)
}

func GetIOSprofileHandler(c *gin.Context) {

	c.FileAttachment("./profile/Hotspot.mobileconfig", "Hotspot.mobileconfig")
}
