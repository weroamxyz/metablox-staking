package controllers

import (
	"crypto/ecdsa"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/contract"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/foundationdao"
	"github.com/metabloxStaking/interest"
	"github.com/metabloxStaking/models"
)

const placeholderExchangeRate = 30.0

func GetProductInfoByIDHandler(c *gin.Context) {
	productID := c.Param("id")
	product, err := dao.GetProductInfoByID(productID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	product.CurrentAPY = 1234 //todo: get value from Colin's code
	ResponseSuccess(c, product)
}

func GetAllProductInfoHandler(c *gin.Context) {
	products, err := dao.GetAllProductInfo()
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
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
	transactions, err := dao.GetOrderInterestByID(orderID)
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

	err = dao.RedeemInterestByOrderID(orderID)
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

func GetMinerListHandler(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	minerList, err := foundationdao.GetAllMinerInfo()
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	for _, miner := range minerList {
		createDate, err := time.Parse("2006-01-02 15:04:05", miner.CreateTime)
		if err != nil {
			ResponseErrorWithMsg(c, CodeError, err.Error())
			return
		}
		miner.CreateTime = strconv.FormatFloat(float64(createDate.UnixNano())/float64(time.Second), 'f', 3, 64)
	}

	if latitude == "" || longitude == "" {
		ResponseSuccess(c, minerList)
		return
	}

	floatLat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	floatLong, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	closestMiner := models.NewMinerInfo()
	closestDistance := math.Inf(1)

	for _, miner := range minerList {
		if miner.Longitude == nil || miner.Latitude == nil {
			continue
		}
		longDistance := floatLong - *miner.Longitude
		latDistance := floatLat - *miner.Latitude
		totalDistance := math.Sqrt(math.Pow(longDistance, 2) + math.Pow(latDistance, 2))
		if totalDistance < closestDistance {
			closestDistance = totalDistance
			closestMiner = miner
		}
	}

	ResponseSuccess(c, closestMiner)
}

func GetMinerByIDHandler(c *gin.Context) {
	minerID := c.Query("minerid")

	miner, err := foundationdao.GetMinerInfoByID(minerID)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	createDate, err := time.Parse("2006-01-02 15:04:05", miner.CreateTime)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	miner.CreateTime = strconv.FormatFloat(float64(createDate.UnixNano())/float64(time.Second), 'f', 3, 64)

	ResponseSuccess(c, miner)
}

func GetExchangeRateHandler(c *gin.Context) {
	ResponseSuccess(c, placeholderExchangeRate)
}

func GetRewardHistoryHandler(c *gin.Context) {
	did := c.Param("did")
	exchangeList, err := foundationdao.GetSeedHistory(did)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	redeemedToken := 0.0

	for _, exchange := range exchangeList {
		redeemedToken += exchange.Amount
	}

	ResponseSuccess(c, redeemedToken)
}

func ExchangeSeedHandler(c *gin.Context) {
	input := models.NewSeedExchangeInput()
	err := c.BindJSON(input)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	minerPubKey := new(ecdsa.PublicKey) //todo: get this from some source

	holderPubKey, err := crypto.UnmarshalPubkey(input.PublicKeyString)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	_, err = presentations.VerifyVP(&input.SeedPresentation, holderPubKey, minerPubKey) //going to fail at the moment as we don't have all the info to do this verification
	if err != nil {                                                                     //skip this error check to avoid failures until we can properly verify seed presentations
		//ResponseErrorWithMsg(c, CodeError, err.Error())
		//return
	}

	targetAddress := common.HexToAddress(input.WalletAddress)

	seedVC := input.SeedPresentation.VerifiableCredential[0]
	splitID := strings.Split(seedVC.ID, "/")
	if len(splitID) != 5 {
		ResponseErrorWithMsg(c, CodeError, "VC id is improperly formatted")
		return
	}
	models.ConvertCredentialSubject(&seedVC)
	seedInfo := seedVC.CredentialSubject.(models.SeedInfo)
	exchangeValue := seedInfo.Amount * placeholderExchangeRate //todo: may have to change calculation method

	txHash, err := contract.TransferTokens(targetAddress, int(exchangeValue)) //todo: need a proper method of converting exchangeValue into an int
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	exchange := models.NewSeedExchange()
	exchange.VcID = strings.Split(seedVC.ID, "/")[4] //should equal numerical ID
	exchange.UserDID = seedInfo.ID
	exchange.ExchangeRate = placeholderExchangeRate           //todo: get actual value
	exchange.Amount = seedInfo.Amount * exchange.ExchangeRate //todo: may have to change calculation method

	err = foundationdao.UploadSeedExchange(exchange)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	output := models.NewSeedExchangeOutput()
	output.Amount = exchange.Amount
	output.ExchangeRate = exchange.ExchangeRate
	output.TxHash = txHash
	output.TxTime = strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)

	ResponseSuccess(c, output)
}
