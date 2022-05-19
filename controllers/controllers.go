package controllers

import (
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/interest"
	"github.com/metabloxStaking/models"
)

func GetProductInfoByIDHandler(c *gin.Context) {
	productID := c.Param("id")
	product, err := dao.GetProductInfoByID(productID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccess(c, product)
}

func GetAllProductInfoHandler(c *gin.Context) {
	products, err := dao.GetAllProductInfo()
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccessWithMsg(c, products)
}

func PurchaseProductByIDHandler(c *gin.Context) {

	ResponseSuccessWithMsg(c, "placeholder")
}

func GetStakingRecordsHandler(c *gin.Context) {
	userDID := c.Param("did")
	records, err := dao.GetStakingRecords(userDID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	stmt, err := dao.PrepareGetInterestByOrderID()
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	for _, record := range records {
		createTime, err := time.Parse("2006-01-02 15:04:05", record.CreateDate)
		if err != nil {
			ResponseErrorWithMsg(c, CodeError, err.Error())
			stmt.Close()
			return
		}
		timeElapsed := time.Since(createTime)
		record.RedeemAll = (timeElapsed.Hours() < 24)
		if record.Type == true {
			record.TotalInterestGain = interest.CalculateTotalInterest(record.OrderID)
			daysElapsed := timeElapsed.Hours() / 24
			record.Term = new(int)
			*record.Term = int(1 + math.Floor(daysElapsed/180))
		} else {
			record.TotalInterestGain, err = dao.ExecuteGetInterestStmt(record.OrderID, stmt)
			if err != nil {
				ResponseErrorWithMsg(c, CodeError, err.Error())
				stmt.Close()
				return
			}
		}
	}
	stmt.Close()
	ResponseSuccessWithMsg(c, records)
}

func GetTransactionsByOrderIDHandler(c *gin.Context) {
	orderID := c.Param("id")
	transactions, err := dao.GetTransactionsByOrderID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccessWithMsg(c, transactions)
}

func RedeemOrderHandler(c *gin.Context) {
	input := models.NewRedeemInput()
	err := c.BindJSON(input)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	createDate, err := dao.GetOrderCreateDate(input.OrderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	createTime, err := time.Parse("2006-01-02 15:04:05", createDate)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	elapsedTime := time.Since(createTime)
	elapsedDays := int(math.Floor(elapsedTime.Hours() / 24))
	daysLeftInTerm := 180 - (elapsedDays % 180)
	if daysLeftInTerm != 1 {
		ResponseErrorWithMsg(c, CodeError, "Order can only be redeemed on final day of term")
		return
	}

	err = dao.RedeemOrderByOrderID(input.OrderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	order, err := dao.GetOrderByID(input.OrderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	txInfo := models.NewTXInfo()
	txInfo.OrderID, _ = strconv.Atoi(input.OrderID)
	txInfo.UserDID = order.UserDID
	txInfo.TXCurrencyType = input.TXCurrencyType
	txInfo.TXType = "Redeem"
	txInfo.Amount = input.Amount
	txInfo.UserAddress = input.UserAddress

	err = dao.UploadTransaction(txInfo)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "Order successfully redeemed!")
}

func RedeemInterestHandler(c *gin.Context) {
	input := models.NewRedeemInput()
	err := c.BindJSON(input)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	valid, err := dao.CheckIfOrderMeetsMinimumInterest(input.OrderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	if !valid {
		ResponseErrorWithMsg(c, CodeError, "Order does not meet minimum interest required to redeem")
		return
	}

	err = dao.RedeemInterestByOrderID(input.OrderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	order, err := dao.GetOrderByID(input.OrderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	txInfo := models.NewTXInfo()
	txInfo.OrderID, _ = strconv.Atoi(input.OrderID)
	txInfo.UserDID = order.UserDID
	txInfo.TXCurrencyType = input.TXCurrencyType
	txInfo.TXType = "Harvest"
	txInfo.Amount = input.Amount
	txInfo.UserAddress = input.UserAddress

	err = dao.UploadTransaction(txInfo)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "Interest successfully redeemed!")
}
