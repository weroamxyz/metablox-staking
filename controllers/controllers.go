package controllers

import (
	"database/sql"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/contract"
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
	var totalPrincipal float64
	update, err := dao.GetLatestPrincipalUpdate(productID)
	if err == nil {
		totalPrincipal = update.TotalPrincipal
	} else if err == sql.ErrNoRows {
		totalPrincipal = 0.0
	} else {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	product.CurrentAPY = interest.CalculateCurrentAPY(product, totalPrincipal)

	ResponseSuccess(c, product)
}

func GetAllProductInfoHandler(c *gin.Context) {
	products, err := dao.GetAllProductInfo()
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	var totalPrincipal float64
	for _, product := range products {
		update, err := dao.GetLatestPrincipalUpdate(product.ID)
		if err == nil {
			totalPrincipal = update.TotalPrincipal
		} else if err == sql.ErrNoRows {
			totalPrincipal = 0.0
		} else {
			ResponseErrorWithMsg(c, CodeError, err.Error())
			return
		}
		product.CurrentAPY = interest.CalculateCurrentAPY(product, totalPrincipal)
	}
	ResponseSuccess(c, products)
}

func CreateOrderHandler(c *gin.Context) {
	var err error
	input := models.NewCreateOrderInput()
	c.BindJSON(input)

	newOrder := models.NewOrder()
	newOrder.ProductID = input.ProductID

	newOrder.UserDID = input.UserDID
	newOrder.Type = models.OrderTypePending
	newOrder.PaymentAddress = "placeholder" //todo: find a way to lookup the correct value from the PaymentInfo table
	newOrder.Term = new(int)
	*newOrder.Term = 1
	newOrder.Amount = input.Amount
	newOrder.UserAddress = input.UserAddress

	orderID, err := dao.CreateOrder(newOrder)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	output := models.NewCreateOrderOutput()

	output.OrderID = strconv.Itoa(orderID)
	output.PaymentAddress = newOrder.PaymentAddress

	ResponseSuccess(c, output)
}

func SubmitBuyinHandler(c *gin.Context) {
	input := models.NewSubmitBuyinInput()
	c.BindJSON(input)

	exists, err := dao.CheckIfTXExists(input.TxHash)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	if exists {
		ResponseErrorWithMsg(c, CodeError, "provided tx hash is already recorded in db")
		return
	}

	completed, err := contract.CheckIfTransactionCompleted(input.TxHash)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	if !completed {
		ResponseErrorWithMsg(c, CodeError, "transaction not yet completed")
		return
	}

	txInfo := models.NewTXInfo()
	txInfo.OrderID = input.OrderID

	order, err := dao.GetOrderByID(input.OrderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	product, err := dao.GetProductInfoByID(order.ProductID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	txInfo.TXCurrencyType = "MBLX"
	txInfo.TXType = "BuyIn"
	txInfo.TXHash = new(string)
	*txInfo.TXHash = input.TxHash
	txInfo.Principal = order.Amount
	txInfo.Interest = 0
	txInfo.UserAddress = order.UserAddress
	txInfo.RedeemableTime = time.Now().AddDate(0, 0, 179).Truncate(24 * time.Hour).Format("2006-01-02 15:04:05.000")
	err = dao.SubmitBuyin(txInfo)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	date, err := dao.GetTXCreateDate(input.TxHash)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	output := models.NewSubmitBuyinOutput()
	output.ProductName = product.ProductName
	output.Amount = order.Amount
	output.Time = date
	output.TXCurrencyType = txInfo.TXCurrencyType
	output.UserAddress = txInfo.UserAddress

	// record change in staking pool's total principal
	newPrincipal := models.NewPrincipalUpdate()
	oldPrincipal, err := dao.GetLatestPrincipalUpdate(product.ID)
	if err == nil {
		newPrincipal.TotalPrincipal = oldPrincipal.TotalPrincipal + txInfo.Principal
	} else if err == sql.ErrNoRows {
		newPrincipal.TotalPrincipal = txInfo.Principal
	} else {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	err = dao.InsertPrincipalUpdate(product.ID, newPrincipal.TotalPrincipal)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
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

		purchaseTime, err := time.Parse("2006-01-02 15:04:05", record.PurchaseTime)
		if err != nil {
			ResponseErrorWithMsg(c, CodeError, err.Error())
			stmt.Close()
			return
		}
		record.PurchaseTime = strconv.FormatFloat(float64(purchaseTime.UnixNano())/float64(time.Second), 'f', 3, 64)

		redeemDate, err := time.Parse("2006-01-02 15:04:05", record.RedeemableTime)
		if err != nil {
			ResponseErrorWithMsg(c, CodeError, err.Error())
			stmt.Close()
			return
		}
		record.RedeemableTime = strconv.FormatFloat(float64(redeemDate.UnixNano())/float64(time.Second), 'f', 3, 64)

		timeElapsed := time.Since(redeemDate)
		record.IsInClosureWindow = (0 < timeElapsed.Hours() && timeElapsed.Hours() < 24)

		if record.OrderStatus == models.OrderTypeHolding {
			interest.CalculateInterest() //query Colin's code to update interest value in db
		}

		interestInfo, err := dao.ExecuteGetInterestStmt(record.OrderID, stmt)
		if err != nil {
			ResponseErrorWithMsg(c, CodeError, err.Error())
			stmt.Close()
			return
		}
		record.InterestGain = interestInfo.AccumulatedInterest - interestInfo.TotalInterestGained
		record.TotalAmount = record.InterestGain + record.PrincipalAmount
	}
	stmt.Close()
	ResponseSuccess(c, records)
}

func GetTransactionsByOrderIDHandler(c *gin.Context) {
	orderID := c.Param("id")
	transactions, err := dao.GetTransactionsByOrderID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, transactions)
}

func GetTransactionsByUserDIDHandler(c *gin.Context) {
	userDID := c.Param("did")
	transactions, err := dao.GetTransactionsByUserDID(userDID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, transactions)
}

func GetOrderInterestHandler(c *gin.Context) {
	orderID := c.Param("id")
	transactions, err := interest.GetOrderInterestList(orderID, time.Now())
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, transactions)
}

func RedeemOrderHandler(c *gin.Context) {
	orderID := c.Param("id")

	redeemableDate, err := dao.GetOrderRedeemableDate(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	redeemableTime, err := time.Parse("2006-01-02 15:04:05", redeemableDate)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	elapsedTime := time.Since(redeemableTime)
	elapsedDays := int(math.Floor(elapsedTime.Hours() / 24))
	if elapsedDays != 0 {
		ResponseErrorWithMsg(c, CodeError, "Order can only be redeemed on final day of term")
		return
	}

	userAddress, err := dao.GetUserAddressByOrderID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	interestInfo, err := dao.GetInterestInfoByOrderID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	productName, err := dao.GetProductNameForOrder(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	order, err := dao.GetOrderByID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	txHash := contract.RedeemOrder()

	txInfo := models.NewTXInfo()
	txInfo.OrderID = orderID
	txInfo.TXCurrencyType = "MBLX"
	txInfo.TXType = "Redeem"
	txInfo.UserAddress = userAddress
	txInfo.RedeemableTime = redeemableDate
	txInfo.TXHash = new(string)
	*txInfo.TXHash = txHash

	err = dao.UploadTransaction(txInfo)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	output := models.NewRedeemOrderOutput()
	output.Amount = (interestInfo.AccumulatedInterest - interestInfo.TotalInterestGained) + order.Amount
	output.ProductName = productName
	output.TXCurrencyType = "MBLX"
	output.TXHash = txHash
	output.Time = strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output.ToAddress = userAddress

	// record change in staking pool's total principal
	newPrincipal := models.NewPrincipalUpdate()
	oldPrincipal, err := dao.GetLatestPrincipalUpdate(order.ProductID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	newPrincipal.TotalPrincipal = oldPrincipal.TotalPrincipal - order.Amount

	err = dao.InsertPrincipalUpdate(order.ProductID, newPrincipal.TotalPrincipal)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
}

func RedeemInterestHandler(c *gin.Context) {
	orderID := c.Param("id")

	minInterest, err := dao.GetMinimumInterestByOrderID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	interestInfo, err := dao.GetInterestInfoByOrderID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	currentInterest := interestInfo.AccumulatedInterest - interestInfo.TotalInterestGained

	if currentInterest < float64(minInterest) {
		ResponseErrorWithMsg(c, CodeError, "order does not meet minimum interest required to redeem")
		return
	}

	txHash := contract.RedeemInterest()

	err = dao.RedeemInterestByOrderID(orderID) // TODO: OrderInterest table doesn't have TotalInterestGain row anymore
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	userAddress, err := dao.GetUserAddressByOrderID(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	txInfo := models.NewTXInfo()
	txInfo.OrderID = orderID
	txInfo.TXCurrencyType = "MBLX"
	txInfo.TXType = "Harvest"
	txInfo.UserAddress = userAddress
	txInfo.RedeemableTime = time.Now().Format("2006-01-02 15:04:05.000")
	txInfo.TXHash = new(string)
	*txInfo.TXHash = txHash

	productName, err := dao.GetProductNameForOrder(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	output := models.NewRedeemOrderOutput()
	output.Amount = currentInterest
	output.ProductName = productName
	output.TXCurrencyType = "MBLX"
	output.TXHash = txHash
	output.Time = strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output.ToAddress = userAddress

	err = dao.UploadTransaction(txInfo)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	err = dao.HarvestOrderInterest(orderID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, output)
}
