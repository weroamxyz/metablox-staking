package dao

import (
	"testing"

	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetProductInfoByID(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	dbProduct, err := GetProductInfoByID("1")
	assert.Nil(t, err)
	assert.Equal(t, "1", dbProduct.ID)
	assert.Equal(t, "TestProduct1", dbProduct.ProductName)
	assert.Equal(t, "10000000", dbProduct.MinOrderValue.String())
	assert.Equal(t, "500000000000", dbProduct.TopUpLimit.String())
	assert.Equal(t, "5000000", dbProduct.MinRedeemValue.String())
	assert.Equal(t, 180, dbProduct.LockUpPeriod)
	assert.Equal(t, 0.2, dbProduct.DefaultAPY)
	assert.Equal(t, "2022-12-27 00:00:00.000", dbProduct.CreateDate)
	assert.Equal(t, "2022-12-27 00:00:00.000", dbProduct.StartDate)
	assert.Equal(t, 1, dbProduct.Term)
	assert.Equal(t, "0", dbProduct.BurnedInterest.String())
	assert.True(t, dbProduct.Status)
	assert.Equal(t, "testPaymentAddress", dbProduct.PaymentAddress)
	assert.Equal(t, "MBLX", dbProduct.CurrencyType)
	assert.Equal(t, "Ethereum", dbProduct.Network)
	assert.Nil(t, dbProduct.NextProductID)
}

func TestGetPaymentAddress(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	dbAddress, err := GetPaymentAddress("1")
	assert.Nil(t, err)
	assert.Equal(t, "testPaymentAddress", dbAddress)

	dbAddress, err = GetPaymentAddress("2")
	assert.Nil(t, err)
	assert.Equal(t, "testPaymentAddress2", dbAddress)
}

func TestGetAllProductInfo(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	dbProducts, err := GetAllProductInfo()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(dbProducts))
	dbProduct := dbProducts[0]
	assert.Equal(t, "1", dbProduct.ID)
	assert.Equal(t, "TestProduct1", dbProduct.ProductName)
	assert.Equal(t, "10000000", dbProduct.MinOrderValue.String())
	assert.Equal(t, "500000000000", dbProduct.TopUpLimit.String())
	assert.Equal(t, "5000000", dbProduct.MinRedeemValue.String())
	assert.Equal(t, 180, dbProduct.LockUpPeriod)
	assert.Equal(t, 0.2, dbProduct.DefaultAPY)
	assert.Equal(t, "2022-12-27 00:00:00.000", dbProduct.CreateDate)
	assert.Equal(t, "2022-12-27 00:00:00.000", dbProduct.StartDate)
	assert.Equal(t, 1, dbProduct.Term)
	assert.Equal(t, "0", dbProduct.BurnedInterest.String())
	assert.True(t, dbProduct.Status)
	assert.Equal(t, "testPaymentAddress", dbProduct.PaymentAddress)
	assert.Equal(t, "MBLX", dbProduct.CurrencyType)
	assert.Equal(t, "Ethereum", dbProduct.Network)
	assert.Nil(t, dbProduct.NextProductID)

	dbProduct = dbProducts[1]
	assert.Equal(t, "2", dbProduct.ID)
	assert.Equal(t, "TestProduct2", dbProduct.ProductName)
	assert.Equal(t, "10000000", dbProduct.MinOrderValue.String())
	assert.Equal(t, "500000000000", dbProduct.TopUpLimit.String())
	assert.Equal(t, "4000000", dbProduct.MinRedeemValue.String())
	assert.Equal(t, 180, dbProduct.LockUpPeriod)
	assert.Equal(t, 0.2, dbProduct.DefaultAPY)
	assert.Equal(t, "2022-06-30 00:00:00.000", dbProduct.CreateDate)
	assert.Equal(t, "2022-06-30 00:00:00.000", dbProduct.StartDate)
	assert.Equal(t, 1, dbProduct.Term)
	assert.Equal(t, "0", dbProduct.BurnedInterest.String())
	assert.True(t, dbProduct.Status)
	assert.Equal(t, "testPaymentAddress2", dbProduct.PaymentAddress)
	assert.Equal(t, "MBLX", dbProduct.CurrencyType)
	assert.Equal(t, "Ethereum", dbProduct.Network)
	assert.Nil(t, dbProduct.NextProductID)

	dbProduct = dbProducts[2]
	assert.Equal(t, "3", dbProduct.ID)
	assert.Equal(t, "TestProduct3", dbProduct.ProductName)
	assert.Equal(t, "10000000", dbProduct.MinOrderValue.String())
	assert.Equal(t, "500000000000", dbProduct.TopUpLimit.String())
	assert.Equal(t, "5000000", dbProduct.MinRedeemValue.String())
	assert.Equal(t, 180, dbProduct.LockUpPeriod)
	assert.Equal(t, 0.2, dbProduct.DefaultAPY)
	assert.Equal(t, "2022-12-27 00:00:00.000", dbProduct.CreateDate)
	assert.Equal(t, "2022-12-27 00:00:00.000", dbProduct.StartDate)
	assert.Equal(t, 1, dbProduct.Term)
	assert.Equal(t, "0", dbProduct.BurnedInterest.String())
	assert.True(t, dbProduct.Status)
	assert.Equal(t, "testPaymentAddress3", dbProduct.PaymentAddress)
	assert.Equal(t, "MBLX", dbProduct.CurrencyType)
	assert.Equal(t, "Ethereum", dbProduct.Network)
	assert.Nil(t, dbProduct.NextProductID)
}

func TestCreateAndGetOrder(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	order := models.NewOrder("1", "did:metablox:sampleDID", models.OrderTypePending, "sampleAddress", decimal.NewFromInt(567), "sampleUserAddress")
	orderID, err := CreateOrder(order)
	assert.Nil(t, err)
	assert.Equal(t, 6, orderID)

	dbOrder, err := GetOrderByID("6")
	assert.Nil(t, err)
	assert.Equal(t, order.AccumulatedInterest, dbOrder.AccumulatedInterest)
	assert.Equal(t, order.Amount, dbOrder.Amount)
	assert.Equal(t, "6", dbOrder.OrderID)
	assert.Equal(t, order.PaymentAddress, dbOrder.PaymentAddress)
	assert.Equal(t, order.ProductID, dbOrder.ProductID)
	assert.Equal(t, order.Term, dbOrder.Term)
	assert.Equal(t, order.TotalInterestGained, dbOrder.TotalInterestGained)
	assert.Equal(t, order.Type, dbOrder.Type)
	assert.Equal(t, order.UserAddress, dbOrder.UserAddress)
	assert.Equal(t, order.UserDID, dbOrder.UserDID)

	order2 := models.NewOrder("1", "did:metablox:sampleDID2", models.OrderTypePending, "sampleAddress2", decimal.NewFromInt(567), "sampleUserAddress2")
	orderID2, err := CreateOrder(order2)
	assert.Nil(t, err)
	assert.Equal(t, 7, orderID2)

	dbOrder2, err := GetOrderByID("7")
	assert.Nil(t, err)
	assert.Equal(t, "7", dbOrder2.OrderID)
}

func TestCheckIfTXExists(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	exists, err := CheckIfTXExists("placeholderHash")
	assert.Nil(t, err)
	assert.True(t, exists)
	exists, err = CheckIfTXExists("badHash")
	assert.Nil(t, err)
	assert.False(t, exists)
}

func TestGetTXCreateDate(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	createDate, err := GetTXCreateDate("placeholderHash")
	assert.Nil(t, err)
	assert.Equal(t, "1653683120.955", createDate)
}

func TestGetStakingRecords(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	records, err := GetStakingRecords("did:metablox:test")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(records))
	record := records[0]
	assert.Equal(t, "1", record.OrderID)
	assert.Equal(t, "1", record.ProductID)
	assert.Equal(t, "Holding", record.OrderStatus)
	assert.Equal(t, 1, *record.Term)
	assert.Equal(t, "2022-05-27 13:25:20.955", record.PurchaseTime)
	assert.Equal(t, "2022-11-21 16:00:00.000", record.RedeemableTime)
	assert.Equal(t, "100", record.PrincipalAmount.String())
	assert.Equal(t, "MBLX", record.TXCurrencyType)
}

func TestGetInterestInfoByOrderID(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	interestInfo, err := GetInterestInfoByOrderID("1")
	assert.Nil(t, err)
	assert.Equal(t, "100", interestInfo.AccumulatedInterest.String())
	assert.Equal(t, "50", interestInfo.TotalInterestGained.String())
}

func TestGetTransactionsByOrderID(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	txs, err := GetTransactionsByOrderID("1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(txs))
	tx1 := txs[0]
	assert.Equal(t, "1", tx1.PaymentNo)
	assert.Equal(t, "1", tx1.OrderID)
	assert.Equal(t, "MBLX", tx1.TXCurrencyType)
	assert.Equal(t, "BuyIn", tx1.TXType)
	assert.Equal(t, "placeholderHash", tx1.TXHash)
	assert.Equal(t, "30", tx1.Principal.String())
	assert.Equal(t, "50", tx1.Interest.String())
	assert.Equal(t, "placeholderUserAddress", tx1.UserAddress)
	assert.Equal(t, "2022-05-27 13:25:20.955", tx1.CreateDate)
	assert.Equal(t, "2022-11-21 16:00:00.000", tx1.RedeemableTime)

	tx2 := txs[1]
	assert.Equal(t, "3", tx2.PaymentNo)
	assert.Equal(t, "1", tx2.OrderID)
	assert.Equal(t, "MBLX", tx2.TXCurrencyType)
	assert.Equal(t, "InterestOnly", tx2.TXType)
	assert.Equal(t, "placeholderHash3", tx2.TXHash)
	assert.Equal(t, "30", tx2.Principal.String())
	assert.Equal(t, "50", tx2.Interest.String())
	assert.Equal(t, "placeholderUserAddress", tx2.UserAddress)
	assert.Equal(t, "2022-05-27 13:25:20.955", tx2.CreateDate)
	assert.Equal(t, "2022-11-21 16:00:00.000", tx2.RedeemableTime)
}

func TestRedeemOrder(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	sampleTX := models.NewTXInfo("3", "MBLX", "OrderClosure", "placeholderHash", decimal.NewFromInt(30), decimal.NewFromInt(50), "placeholderAddress", "2022-06-30 00:00:00.000")
	err = RedeemOrder(sampleTX)
	assert.Nil(t, err)
	dbOrder, err := GetOrderByID("3")
	assert.Nil(t, err)
	assert.Equal(t, "Complete", dbOrder.Type)
	assert.Equal(t, dbOrder.AccumulatedInterest, dbOrder.TotalInterestGained)
	dbTxs, err := GetTransactionsByOrderID("3")
	assert.Nil(t, err)
	dbTX := dbTxs[1]
	assert.Equal(t, sampleTX.Interest, dbTX.Interest)
	assert.Equal(t, sampleTX.OrderID, dbTX.OrderID)
	assert.Equal(t, "5", dbTX.PaymentNo)
	assert.Equal(t, sampleTX.Principal, dbTX.Principal)
	assert.Equal(t, sampleTX.RedeemableTime, dbTX.RedeemableTime)
	assert.Equal(t, sampleTX.TXCurrencyType, dbTX.TXCurrencyType)
	assert.Equal(t, sampleTX.TXHash, dbTX.TXHash)
	assert.Equal(t, sampleTX.TXType, dbTX.TXType)
	assert.Equal(t, sampleTX.UserAddress, dbTX.UserAddress)
}

func TestRedeemInterest(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	sampleTX := models.NewTXInfo("3", "MBLX", "OrderClosure", "placeholderHash", decimal.NewFromInt(30), decimal.NewFromInt(50), "placeholderAddress", "2022-06-30 00:00:00.000")
	err = RedeemInterest(sampleTX)
	assert.Nil(t, err)
	dbOrder, err := GetOrderByID("3")
	assert.Nil(t, err)
	assert.Equal(t, "Holding", dbOrder.Type)
	assert.Equal(t, dbOrder.AccumulatedInterest, dbOrder.TotalInterestGained)
	dbTxs, err := GetTransactionsByOrderID("3")
	assert.Nil(t, err)
	dbTX := dbTxs[1]
	assert.Equal(t, sampleTX.Interest, dbTX.Interest)
	assert.Equal(t, sampleTX.OrderID, dbTX.OrderID)
	assert.Equal(t, "5", dbTX.PaymentNo)
	assert.Equal(t, sampleTX.Principal, dbTX.Principal)
	assert.Equal(t, sampleTX.RedeemableTime, dbTX.RedeemableTime)
	assert.Equal(t, sampleTX.TXCurrencyType, dbTX.TXCurrencyType)
	assert.Equal(t, sampleTX.TXHash, dbTX.TXHash)
	assert.Equal(t, sampleTX.TXType, dbTX.TXType)
	assert.Equal(t, sampleTX.UserAddress, dbTX.UserAddress)
}

func TestGetHoldingOrderIDsForProduct(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	ids, err := GetHoldingOrderIDsForProduct("1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(ids))
	assert.Equal(t, "1", ids[0])
	assert.Equal(t, "2", ids[1])
	ids, err = GetHoldingOrderIDsForProduct("2")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(ids))
	assert.Equal(t, "5", ids[0])
}

func TestGetOrdersByProductID(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	orders, err := GetOrdersByProductID("1")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(orders))
	assert.Equal(t, "1", orders[0].OrderID)
	assert.Equal(t, "placeholder", orders[0].PaymentAddress)
	assert.Equal(t, "2", orders[1].OrderID)
	assert.Equal(t, "placeholder2", orders[1].PaymentAddress)
	assert.Equal(t, "3", orders[2].OrderID)
	assert.Equal(t, "placeholder3", orders[2].PaymentAddress)
	orders, err = GetOrdersByProductID("2")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, "5", orders[0].OrderID)
	assert.Equal(t, "placeholder5", orders[0].PaymentAddress)
}

func TestGetOrderCreateDate(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	createDate, err := GetOrderCreateDate("1")
	assert.Nil(t, err)
	assert.Equal(t, "2022-05-27 13:25:20.955", createDate)
	createDate, err = GetOrderCreateDate("2")
	assert.Nil(t, err)
	assert.Equal(t, "2022-05-27 13:25:20.955", createDate)
	_, err = GetOrderCreateDate("3")
	assert.NotNil(t, err)
}

func TestGetOrderRedeemableDate(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	createDate, err := GetOrderRedeemableDate("1")
	assert.Nil(t, err)
	assert.Equal(t, "2022-11-21 16:00:00.000", createDate)
	createDate, err = GetOrderRedeemableDate("2")
	assert.Nil(t, err)
	assert.Equal(t, "2022-11-21 16:00:00.000", createDate)
	_, err = GetOrderRedeemableDate("3")
	assert.NotNil(t, err)
}

func TestGetUserAddressByOrderID(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	address, err := GetUserAddressByOrderID("1")
	assert.Nil(t, err)
	assert.Equal(t, "placeholderUserAddress", address)
	address, err = GetUserAddressByOrderID("3")
	assert.Nil(t, err)
	assert.Equal(t, "placeholderUserAddress3", address)
}

func TestCompareMinimumInterest(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	meetsMinInterest, err := CompareMinimumInterest("1", "5000000")
	assert.Nil(t, err)
	assert.True(t, meetsMinInterest)
	meetsMinInterest, err = CompareMinimumInterest("1", "4999999")
	assert.Nil(t, err)
	assert.False(t, meetsMinInterest)
	meetsMinInterest, err = CompareMinimumInterest("5", "4999999")
	assert.Nil(t, err)
	assert.True(t, meetsMinInterest)
}

func TestUploadTransaction(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	sampleTX := models.NewTXInfo("5", "MBLX", "OrderClosure", "placeholderHash", decimal.NewFromInt(30), decimal.NewFromInt(50), "placeholderAddress", "2022-06-30 00:00:00.000")
	err = UploadTransaction(sampleTX)
	assert.Nil(t, err)
	dbTXs, err := GetTransactionsByOrderID("5")
	assert.Nil(t, err)
	dbTX := dbTXs[0]
	assert.Equal(t, sampleTX.Interest, dbTX.Interest)
	assert.Equal(t, sampleTX.OrderID, dbTX.OrderID)
	assert.Equal(t, "5", dbTX.PaymentNo)
	assert.Equal(t, sampleTX.Principal, dbTX.Principal)
	assert.Equal(t, sampleTX.RedeemableTime, dbTX.RedeemableTime)
	assert.Equal(t, sampleTX.TXCurrencyType, dbTX.TXCurrencyType)
	assert.Equal(t, sampleTX.TXHash, dbTX.TXHash)
	assert.Equal(t, sampleTX.TXType, dbTX.TXType)
	assert.Equal(t, sampleTX.UserAddress, dbTX.UserAddress)
}

func TestSubmitBuyin(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	sampleTX := models.NewTXInfo("4", "MBLX", "BuyIn", "placeholderHash", decimal.NewFromInt(30), decimal.NewFromInt(50), "placeholderAddress", "2022-06-30 00:00:00.000")
	err = SubmitBuyin(sampleTX)
	assert.Nil(t, err)
	order, err := GetOrderByID("4")
	assert.Nil(t, err)
	assert.Equal(t, "Holding", order.Type)
	dbTXs, err := GetTransactionsByOrderID("4")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(dbTXs))
	dbTX := dbTXs[0]
	assert.Equal(t, sampleTX.Interest, dbTX.Interest)
	assert.Equal(t, sampleTX.OrderID, dbTX.OrderID)
	assert.Equal(t, "5", dbTX.PaymentNo)
	assert.Equal(t, sampleTX.Principal, dbTX.Principal)
	assert.Equal(t, sampleTX.RedeemableTime, dbTX.RedeemableTime)
	assert.Equal(t, sampleTX.TXCurrencyType, dbTX.TXCurrencyType)
	assert.Equal(t, sampleTX.TXHash, dbTX.TXHash)
	assert.Equal(t, sampleTX.TXType, dbTX.TXType)
	assert.Equal(t, sampleTX.UserAddress, dbTX.UserAddress)
}

func TestGetTotalInterestGained(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	dbInterest, err := GetTotalInterestGained("1")
	assert.Nil(t, err)
	assert.Equal(t, "50", dbInterest.String())
}

func TestHarvestOrderInterest(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	err = HarvestOrderInterest("1")
	assert.Nil(t, err)
	dbInterest, err := GetTotalInterestGained("1")
	assert.Nil(t, err)
	assert.Equal(t, "100", dbInterest.String())
}

func TestGetProductNameForOrder(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	name, err := GetProductNameForOrder("1")
	assert.Nil(t, err)
	assert.Equal(t, "TestProduct1", name)
	name, err = GetProductNameForOrder("5")
	assert.Nil(t, err)
	assert.Equal(t, "TestProduct2", name)
}

func TestDisableProduct(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	err = DisableProduct("3")
	assert.Nil(t, err)
	product, err := GetProductInfoByID("3")
	assert.Nil(t, err)
	assert.Equal(t, false, product.Status)
}

func TestInsertAndGetPrincipalUpdate(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	err = InsertPrincipalUpdate("1", "500")
	assert.Nil(t, err)
	err = InsertPrincipalUpdate("1", "100")
	assert.Nil(t, err)
	err = InsertPrincipalUpdate("2", "300")
	assert.Nil(t, err)
	principalUpdates, err := GetPrincipalUpdates("1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(principalUpdates))
	assert.Equal(t, "4", principalUpdates[0].ID)
	assert.Equal(t, "1", principalUpdates[0].ProductID)
	assert.Equal(t, "500", principalUpdates[0].TotalPrincipal.String())
	assert.Equal(t, "5", principalUpdates[1].ID)
	assert.Equal(t, "1", principalUpdates[1].ProductID)
	assert.Equal(t, "100", principalUpdates[1].TotalPrincipal.String())
	principalUpdates, err = GetPrincipalUpdates("2")
	assert.Nil(t, err)
	assert.Equal(t, "6", principalUpdates[0].ID)
	assert.Equal(t, "2", principalUpdates[0].ProductID)
	assert.Equal(t, "300", principalUpdates[0].TotalPrincipal.String())
}

func TestGetLatestPrincipalUpdate(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	latestUpdate, err := GetLatestPrincipalUpdate("3")
	assert.Nil(t, err)
	assert.Equal(t, "888", latestUpdate.TotalPrincipal.String())
}

func TestInsertOrderInterestList(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	sampleInterestList := models.CreateOrderInterestList()
	sampleInterest := models.NewOrderInterest("2", "2022-10-06 00:00:00", 532.1, decimal.NewFromInt(111), decimal.NewFromInt(222))
	sampleInterestList = append(sampleInterestList, sampleInterest)
	sampleInterest = models.NewOrderInterest("2", "2022-11-06 00:00:00", 600.5, decimal.NewFromInt(333), decimal.NewFromInt(444))
	sampleInterestList = append(sampleInterestList, sampleInterest)
	err = InsertOrderInterestList(sampleInterestList)
	assert.Nil(t, err)
	dbInterestList, err := GetOrderInterestByID("2")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(dbInterestList))
	sampleInterest = sampleInterestList[0]
	dbInterest := dbInterestList[0]
	assert.Equal(t, sampleInterest.APY, dbInterest.APY)
	assert.Equal(t, "4", dbInterest.ID)
	assert.Equal(t, sampleInterest.InterestGain, dbInterest.InterestGain)
	assert.Equal(t, sampleInterest.OrderID, dbInterest.OrderID)
	assert.Equal(t, sampleInterest.Time, dbInterest.Time)
	assert.Equal(t, sampleInterest.TotalInterestGain, dbInterest.TotalInterestGain)
	sampleInterest = sampleInterestList[1]
	dbInterest = dbInterestList[1]
	assert.Equal(t, sampleInterest.APY, dbInterest.APY)
	assert.Equal(t, "5", dbInterest.ID)
	assert.Equal(t, sampleInterest.InterestGain, dbInterest.InterestGain)
	assert.Equal(t, sampleInterest.OrderID, dbInterest.OrderID)
	assert.Equal(t, sampleInterest.Time, dbInterest.Time)
	assert.Equal(t, sampleInterest.TotalInterestGain, dbInterest.TotalInterestGain)
}

func TestGetSortedOrderInterestListUntilDate(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	sortedList, err := GetSortedOrderInterestListUntilDate("1", "2022-12-31 00:00:00")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(sortedList))
	assert.Equal(t, "3", sortedList[0].ID)
	assert.Equal(t, "1", sortedList[1].ID)
	assert.Equal(t, "2", sortedList[2].ID)
	sortedList, err = GetSortedOrderInterestListUntilDate("1", "2022-11-31 00:00:00")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(sortedList))
	assert.Equal(t, "3", sortedList[0].ID)
	assert.Equal(t, "1", sortedList[1].ID)
}

func TestUpdateOrderAccumulatedInterest(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	err = UpdateOrderAccumulatedInterest("1", "5000")
	assert.Nil(t, err)
	order, err := GetOrderByID("1")
	assert.Nil(t, err)
	assert.Equal(t, "5000", order.AccumulatedInterest.String())
}

func TestUpdateOrderNewProductID(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)
	err = UpdateOrderNewProductID("1", "4")
	assert.Nil(t, err)
	order, err := GetOrderByID("1")
	assert.Nil(t, err)
	assert.Equal(t, "4", order.ProductID)
}

func TestInsertAndGetMiningRole(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	sampleRole := models.NewMiningRole("did:metablox:test", "testAddress", "testType")
	err = InsertMiningRole(sampleRole)
	assert.Nil(t, err)
	dbRole, err := GetMiningRole("did:metablox:test")
	assert.Nil(t, err)
	assert.Equal(t, sampleRole.DID, dbRole.DID)
	assert.Equal(t, sampleRole.Type, dbRole.Type)
	assert.Equal(t, sampleRole.WalletAddress, dbRole.WalletAddress)
}

func TestGetAllMinerInfo(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	miners, err := GetAllMinerInfo()
	assert.Nil(t, err)
	assert.Equal(t, 4, len(miners))
	miner1 := miners[0]
	assert.Equal(t, "1", miner1.ID)
	assert.Equal(t, "testName", miner1.Name)
	assert.Equal(t, "testSSID", *miner1.SSID)
	assert.Equal(t, "testBSSID", *miner1.BSSID)
	assert.Equal(t, "2022-04-19 00:00:00.000", miner1.CreateTime)
	assert.Equal(t, float64(50), *miner1.Longitude)
	assert.Equal(t, float64(100), *miner1.Latitude)
	assert.True(t, miner1.OnlineStatus)
	assert.Equal(t, float64(25), *miner1.MiningPower)
	assert.True(t, miner1.IsMinable)
	assert.Equal(t, "did:metablox:test", miner1.DID)
	assert.Equal(t, "sampleHost", miner1.Host)
	assert.False(t, miner1.IsVirtual)

	miner3 := miners[2]
	assert.Equal(t, "3", miner3.ID)
	assert.Equal(t, "testName3", miner3.Name)
	assert.Equal(t, "testSSID3", *miner3.SSID)
	assert.Equal(t, "testBSSID3", *miner3.BSSID)
	assert.Equal(t, "2022-04-21 00:00:00.000", miner3.CreateTime)
	assert.Nil(t, miner3.Longitude)
	assert.Nil(t, miner3.Latitude)
	assert.True(t, miner3.OnlineStatus)
	assert.Equal(t, float64(75), *miner3.MiningPower)
	assert.True(t, miner3.IsMinable)
	assert.Equal(t, "did:metablox:test3", miner3.DID)
	assert.Equal(t, "sampleHost", miner3.Host)
	assert.True(t, miner3.IsVirtual)
}

func TestGetAllVirtualMinerInfo(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	miners, err := GetAllVirtualMinerInfo()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(miners))

	miner3 := miners[0]
	assert.Equal(t, "3", miner3.ID)
	assert.Equal(t, "testName3", miner3.Name)
	assert.Equal(t, "testSSID3", *miner3.SSID)
	assert.Equal(t, "testBSSID3", *miner3.BSSID)
	assert.Equal(t, "2022-04-21 00:00:00.000", miner3.CreateTime)
	assert.Nil(t, miner3.Longitude)
	assert.Nil(t, miner3.Latitude)
	assert.True(t, miner3.OnlineStatus)
	assert.Equal(t, float64(75), *miner3.MiningPower)
	assert.True(t, miner3.IsMinable)
	assert.Equal(t, "did:metablox:test3", miner3.DID)
	assert.Equal(t, "sampleHost", miner3.Host)
	assert.True(t, miner3.IsVirtual)

	miner4 := miners[1]
	assert.Equal(t, "4", miner4.ID)
	assert.Equal(t, "testName4", miner4.Name)
	assert.Equal(t, "testSSID4", *miner4.SSID)
	assert.Equal(t, "testBSSID4", *miner4.BSSID)
	assert.Equal(t, "2022-04-22 00:00:00.000", miner4.CreateTime)
	assert.Nil(t, miner4.Longitude)
	assert.Nil(t, miner4.Latitude)
	assert.True(t, miner4.OnlineStatus)
	assert.Equal(t, float64(100), *miner4.MiningPower)
	assert.True(t, miner4.IsMinable)
	assert.Equal(t, "did:metablox:test4", miner4.DID)
	assert.Equal(t, "sampleHost", miner4.Host)
	assert.True(t, miner4.IsVirtual)
}

func TestGetMinerInfoByID(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	miner, err := GetMinerInfoByID("2")
	assert.Nil(t, err)

	assert.Equal(t, "2", miner.ID)
	assert.Equal(t, "testName2", miner.Name)
	assert.Equal(t, "testSSID2", *miner.SSID)
	assert.Equal(t, "testBSSID2", *miner.BSSID)
	assert.Equal(t, "2022-04-20 00:00:00.000", miner.CreateTime)
	assert.Equal(t, float64(75), *miner.Longitude)
	assert.Equal(t, float64(25), *miner.Latitude)
	assert.True(t, miner.OnlineStatus)
	assert.Equal(t, float64(50), *miner.MiningPower)
	assert.True(t, miner.IsMinable)
	assert.Equal(t, "did:metablox:test2", miner.DID)
	assert.Equal(t, "sampleHost", miner.Host)
	assert.False(t, miner.IsVirtual)
}

func TestGetSeedHistory(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	exchanges, err := GetSeedHistory("did:metablox:sampleUser")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(exchanges))
	exchange1 := exchanges[0]
	assert.Equal(t, "did:metablox:sampleUser", exchange1.UserDID)
	assert.Equal(t, "did:metablox:sampleTarget", exchange1.TargetDID)
	assert.Equal(t, "sampleChallenge", exchange1.Challenge)
	assert.Equal(t, "50", exchange1.ExchangeRate)
	assert.Equal(t, "123", exchange1.Amount)
	assert.Equal(t, "2022-04-19 00:00:00.000", exchange1.CreateTime)

	exchange2 := exchanges[1]
	assert.Equal(t, "did:metablox:sampleUser", exchange2.UserDID)
	assert.Equal(t, "did:metablox:sampleTarget2", exchange2.TargetDID)
	assert.Equal(t, "sampleChallenge2", exchange2.Challenge)
	assert.Equal(t, "500", exchange2.ExchangeRate)
	assert.Equal(t, "1234", exchange2.Amount)
	assert.Equal(t, "2022-04-20 00:00:00.000", exchange2.CreateTime)
}

func TestUploadSeedExchange(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	sampleExchange := models.NewSeedExchange("did:metablox:sampleUser3", "did:metablox:sampleTarget3", "sampleChallenge", "123", "321")
	err = UploadSeedExchange(sampleExchange)
	assert.Nil(t, err)

	exchanges, err := GetSeedHistory("did:metablox:sampleUser3")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(exchanges))
	dbExchange := exchanges[0]
	assert.Equal(t, sampleExchange.Amount, dbExchange.Amount)
	assert.Equal(t, sampleExchange.Challenge, dbExchange.Challenge)
	assert.Equal(t, sampleExchange.ExchangeRate, dbExchange.ExchangeRate)
	assert.Equal(t, sampleExchange.TargetDID, dbExchange.TargetDID)
	assert.Equal(t, sampleExchange.UserDID, dbExchange.UserDID)
}

func TestCheckIfExchangeExists(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	keys := models.NewSeedHistoryKeys("did:metablox:sampleUser", "did:metablox:sampleTarget", "sampleChallenge")
	err = CheckIfExchangeExists(keys)
	assert.Equal(t, errval.ErrSeedAlreadyExchanged, err)

	keys = models.NewSeedHistoryKeys("did:metablox:sampleUser5", "did:metablox:sampleTarget", "sampleChallenge")
	err = CheckIfExchangeExists(keys)
	assert.Nil(t, err)

	keys = models.NewSeedHistoryKeys("did:metablox:sampleUser", "did:metablox:sampleTarget5", "sampleChallenge")
	err = CheckIfExchangeExists(keys)
	assert.Nil(t, err)

	keys = models.NewSeedHistoryKeys("did:metablox:sampleUser", "did:metablox:sampleTarget", "sampleChallenge5")
	err = CheckIfExchangeExists(keys)
	assert.Nil(t, err)
}

func TestGetExchangeRate(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	rate, err := GetExchangeRate("1")
	assert.Nil(t, err)
	assert.Equal(t, "721", rate.String())
}

func TestCheckIfDIDIsValidator(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	result, err := CheckIfDIDIsValidator("did:metablox:test")
	assert.Nil(t, err)
	assert.True(t, result)

	result, err = CheckIfDIDIsValidator("did:metablox:test2")
	assert.Nil(t, err)
	assert.False(t, result)

	result, err = CheckIfDIDIsValidator("did:metablox:test3")
	assert.Nil(t, err)
	assert.False(t, result)
}

func TestCheckIfDIDIsMiner(t *testing.T) {
	t.Cleanup(CleanupTestDB)
	err := InitTestDB()
	assert.Nil(t, err)

	result, err := CheckIfDIDIsMiner("did:metablox:test")
	assert.Nil(t, err)
	assert.True(t, result)

	result, err = CheckIfDIDIsMiner("did:metablox:test2")
	assert.Nil(t, err)
	assert.False(t, result)
}
