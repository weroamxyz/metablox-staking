package dao

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
)

var SqlDB *sqlx.DB
var validate *validator.Validate

func InitSql(validatePtr *validator.Validate) error {
	validate = validatePtr

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	SqlDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		logger.Error("Failed to open database: " + err.Error())
		return err
	}

	//Set the maximum number of database connections

	SqlDB.SetConnMaxLifetime(100)

	//Set the maximum number of idle connections on the database

	SqlDB.SetMaxIdleConns(10)

	//Verify connection

	if err := SqlDB.Ping(); err != nil {
		logger.Error("open database fail: ", err)
		return err
	}
	logger.Info("connect success")
	return nil
}

func GetProductInfoByID(productID string) (*models.StakingProduct, error) {
	product := models.CreateStakingProduct()

	sqlStr := "select * from StakingProducts where ID = ?"
	err := SqlDB.Get(product, sqlStr, productID)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func GetPaymentAddress(productID string) (string, error) {
	var address string

	sqlStr := "select PaymentAddress from StakingProducts where ID = ?"
	err := SqlDB.Get(&address, sqlStr, productID)
	if err != nil {
		return "", err
	}

	return address, nil
}

func GetAllProductInfo() ([]*models.StakingProduct, error) {
	var products []*models.StakingProduct
	sqlStr := "select * from StakingProducts"
	rows, err := SqlDB.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.CreateStakingProduct()
		err = rows.StructScan(product)
		if err != nil {
			return nil, err
		}
		err = validate.Struct(product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

func CreateOrder(order *models.Order) (int, error) {
	err := validate.Struct(order)
	if err != nil {
		return 0, err
	}

	sqlStr := "insert into Orders (ProductID, UserDID, Type, Term, PaymentAddress, Amount, UserAddress) values (:ProductID, :UserDID, :Type, :Term, :PaymentAddress, :Amount, :UserAddress)"
	result, err := SqlDB.NamedExec(sqlStr, order)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func CheckIfTXExists(txHash string) (bool, error) {
	var count int
	sqlStr := "select count(*) from TXInfo where TXHash = ?"
	err := SqlDB.Get(&count, sqlStr, txHash)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}

func GetTXCreateDate(txHash string) (string, error) {
	var date string
	sqlStr := "select unix_timestamp(CreateDate) from TXInfo where TXHash = ?"
	err := SqlDB.Get(&date, sqlStr, txHash)
	if err != nil {
		return "", err
	}
	return date, nil
}

func GetStakingRecords(did string) ([]*models.StakingRecord, error) {
	var records []*models.StakingRecord
	sqlStr := "select Orders.OrderID, Orders.ProductID, Orders.Type, Orders.Term, TXInfo.CreateDate, Orders.Amount, TXInfo.TXCurrencyType, TXInfo.RedeemableTime from Orders join TXInfo on TXInfo.OrderID = Orders.OrderID where Orders.UserDID = ? and TXInfo.TXType = 'BuyIn' and Orders.Type != 'Pending'"
	rows, err := SqlDB.Queryx(sqlStr, did)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		record := models.CreateStakingRecord()
		err = rows.StructScan(record)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func GetInterestInfoByOrderID(id string) (*models.OrderInterestInfo, error) {
	info := models.CreateOrderInterestInfo()
	sqlStr := "select AccumulatedInterest, TotalInterestGained from Orders where OrderID = ?"
	err := SqlDB.Get(info, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func PrepareGetInterestByOrderID() (*sqlx.Stmt, error) {
	sqlStr := "select AccumulatedInterest, TotalInterestGained from Orders where OrderID = ?"
	stmt, err := SqlDB.Preparex(sqlStr)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func ExecuteGetInterestStmt(id string, stmt *sqlx.Stmt) (*models.OrderInterestInfo, error) {
	info := models.CreateOrderInterestInfo()
	err := stmt.Get(info, id)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func GetTransactionsByOrderID(orderID string) ([]*models.TXInfo, error) {
	var transactions []*models.TXInfo
	sqlStr := "select PaymentNo, OrderID, TXCurrencyType, TXType, TXHash, Principal, Interest, UserAddress, CreateDate, RedeemableTime from TXInfo where OrderID = ?"
	rows, err := SqlDB.Queryx(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tx := models.CreateTXInfo()
		err = rows.StructScan(tx)
		if err != nil {
			return nil, err
		}
		err = validate.Struct(tx)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, err
}

func GetTransactionsByUserDID(userDID string) ([]*models.TXInfo, error) {
	var transactions []*models.TXInfo
	sqlStr := "select TXInfo.PaymentNo, TXInfo.OrderID, TXInfo.TXCurrencyType, TXInfo.TXType, TXInfo.TXHash, TXInfo.Principal, TXInfo.Interest, TXInfo.UserAddress, TXInfo.CreateDate, TXInfo.RedeemableTime from TXInfo join Orders on Orders.OrderID = TXInfo.OrderID where Orders.UserDID = ?"
	rows, err := SqlDB.Queryx(sqlStr, userDID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tx := models.CreateTXInfo()
		err = rows.StructScan(tx)
		if err != nil {
			return nil, err
		}
		err = validate.Struct(tx)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, err
}

func GetOrderInterestByID(orderID string) ([]*models.OrderInterest, error) {
	var interests []*models.OrderInterest
	sqlStr := "select * from OrderInterest where OrderID = ?"
	rows, err := SqlDB.Queryx(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		interest := models.CreateOrderInterest()
		err = rows.StructScan(interest)
		if err != nil {
			return nil, err
		}
		err = validate.Struct(interest)
		if err != nil {
			return nil, err
		}
		interests = append(interests, interest)
	}
	return interests, nil
}

func RedeemOrder(txInfo *models.TXInfo, interestGained float64) error {
	dbtx, err := SqlDB.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			dbtx.Commit()
		} else {
			dbtx.Rollback()
		}
	}()

	sqlStr := "update OrderInterest set TotalInterestGain = ? where OrderID = ? order by ID desc limit 1"
	_, err = dbtx.Exec(sqlStr, interestGained, txInfo.OrderID)
	if err != nil {
		return err
	}

	sqlStr = "update Orders set TotalInterestGained = AccumulatedInterest, Type = 'Complete' where OrderID = ?"
	_, err = dbtx.Exec(sqlStr, txInfo.OrderID)
	if err != nil {
		return err
	}

	sqlStr = "insert into TXInfo (OrderID, TXCurrencyType, TXType, TXHash, Principal, Interest, UserAddress, RedeemableTime) values (:OrderID, :TXCurrencyType, :TXType, :TXHash, :Principal, :Interest, :UserAddress, :RedeemableTime)"
	_, err = dbtx.NamedExec(sqlStr, txInfo)
	if err != nil {
		return err
	}

	return nil
}

func RedeemInterestByOrderID(orderID string, interestGained float64) error {
	sqlStr := "update OrderInterest set TotalInterestGain = ? where OrderID = ? order by ID desc limit 1"
	_, err := SqlDB.Exec(sqlStr, interestGained, orderID)
	if err != nil {
		return err
	}
	return nil
}

func GetHoldingOrderIDsForProduct(productID string) ([]string, error) {
	var ids []string
	sqlStr := `select distinct Orders.OrderID from Orders where Type = 'Holding' and ProductID = ? and not exists (select * from TXInfo where TXInfo.OrderID = Orders.OrderID and TXInfo.TXType = 'OrderClosure')`
	rows, err := SqlDB.Queryx(sqlStr, productID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		id := new(string)
		err = rows.Scan(id)
		if err != nil {
			logger.Warn(err)
			continue
		}
		ids = append(ids, *id)
	}
	return ids, nil
}

func GetOrdersByProductID(productID string) ([]*models.Order, error) {
	var orders []*models.Order
	sqlStr := `select * from Orders where ProductID = ? and Type = 'Holding'`
	rows, err := SqlDB.Queryx(sqlStr, productID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		order := models.CreateOrder()
		err = rows.StructScan(order)
		if err != nil {
			logger.Warn(err)
			continue
		}
		err = validate.Struct(order)
		if err != nil {
			logger.Warn(err)
			continue
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func GetOrderByID(orderID string) (*models.Order, error) {
	order := models.CreateOrder()
	sqlStr := "select * from Orders where OrderID = ?"
	err := SqlDB.Get(order, sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderCreateDate(orderID string) (string, error) {
	var createDate string
	sqlStr := "select CreateDate from TXInfo where OrderID = ? and TXType = 'BuyIn'"
	err := SqlDB.Get(&createDate, sqlStr, orderID)
	if err != nil {
		return "", err
	}
	return createDate, nil
}

func GetOrderRedeemableDate(orderID string) (string, error) {
	var redeemableDate string
	sqlStr := "select RedeemableTime from TXInfo where OrderID = ? and TXType = 'BuyIn'"
	err := SqlDB.Get(&redeemableDate, sqlStr, orderID)
	if err != nil {
		return "", err
	}
	return redeemableDate, nil
}

func GetUserAddressByOrderID(orderID string) (string, error) {
	var userAddress string
	sqlStr := "select UserAddress from Orders where OrderID = ?"
	err := SqlDB.Get(&userAddress, sqlStr, orderID)
	if err != nil {
		return "", err
	}
	return userAddress, nil
}

func GetOrderBuyInPrincipal(orderID string) (float64, error) {
	var buyInAmount float64
	sqlStr := "select Principal from TXInfo where OrderID = ? and TXType = 'BuyIn'"
	err := SqlDB.Get(&buyInAmount, sqlStr, orderID)
	if err != nil {
		return 0.0, err
	}
	return buyInAmount, nil
}

func CompareMinimumInterest(orderID string, currentInterest float64) (bool, error) {
	var result bool
	sqlStr := "select StakingProducts.MinRedeemValue <= ? from StakingProducts join Orders on StakingProducts.ID = Orders.ProductID where Orders.OrderID = ?"
	err := SqlDB.Get(&result, sqlStr, currentInterest, orderID)
	if err != nil {
		return false, err
	}

	return result, nil
}

func UploadTransaction(tx *models.TXInfo) error {
	err := validate.Struct(tx)
	if err != nil {
		return err
	}
	sqlStr := "insert into TXInfo (OrderID, TXCurrencyType, TXType, TXHash, Principal, Interest, UserAddress, RedeemableTime) values (:OrderID, :TXCurrencyType, :TXType, :TXHash, :Principal, :Interest, :UserAddress, :RedeemableTime)"
	_, err = SqlDB.NamedExec(sqlStr, tx)
	if err != nil {
		return err
	}
	return nil
}

func SubmitBuyin(tx *models.TXInfo) error {
	err := validate.Struct(tx)
	if err != nil {
		return err
	}
	dbTX, err := SqlDB.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			dbTX.Commit()
		} else {
			dbTX.Rollback()
		}
	}()

	sqlStr := "update Orders set Type = 'Holding' where OrderID = ?"
	result, err := dbTX.Exec(sqlStr, tx.OrderID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errval.ErrUpdateOrderStatus
	}

	sqlStr = "insert into TXInfo (OrderID, TXCurrencyType, TXType, TXHash, Principal, Interest, UserAddress, RedeemableTime) values (:OrderID, :TXCurrencyType, :TXType, :TXHash, :Principal, :Interest, :UserAddress, :RedeemableTime)"
	_, err = dbTX.NamedExec(sqlStr, tx)
	if err != nil {
		return err
	}
	return nil
}

func GetTotalInterestGained(id string) (float64, error) {
	var interest float64
	sqlStr := "select TotalInterestGained from Orders where OrderID = ?"
	err := SqlDB.Get(&interest, sqlStr, id)
	if err != nil {
		return 0, err
	}
	return interest, nil
}

func HarvestOrderInterest(id string) error {
	sqlStr := "update Orders set TotalInterestGained = AccumulatedInterest where OrderID = ?"
	_, err := SqlDB.Query(sqlStr, id)
	return err
}

func GetProductNameForOrder(id string) (string, error) {
	var name string
	sqlStr := "select StakingProducts.ProductName from StakingProducts join Orders on StakingProducts.ID = Orders.ProductID where Orders.OrderID = ?"
	err := SqlDB.Get(&name, sqlStr, id)
	if err != nil {
		return "", err
	}
	return name, nil
}

func UpdateProductStatus(id string, status bool) error {
	sqlStr := `update StakingProducts set Status = ? where ID = ?`
	_, err := SqlDB.Exec(sqlStr, status, id)
	return err
}

func InsertPrincipalUpdate(productID string, totalPrincipal float64) error {
	sqlStr := `insert into PrincipalUpdates (ProductID, TotalPrincipal) values (?, ?)`
	_, err := SqlDB.Exec(sqlStr, productID, totalPrincipal)
	return err
}

func GetLatestPrincipalUpdate(productID string) (*models.PrincipalUpdate, error) {
	update := models.NewPrincipalUpdate()

	sqlStr := "select * from PrincipalUpdates where ProductID = ? order by Time desc"
	err := SqlDB.Get(update, sqlStr, productID)
	if err != nil {
		return nil, err
	}

	return update, nil
}

func GetPrincipalUpdates(productID string) ([]*models.PrincipalUpdate, error) {
	sqlStr := `select * from PrincipalUpdates where ProductID = ? order by Time asc`
	rows, err := SqlDB.Queryx(sqlStr, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	updates := models.NewPrincipalUpdateList()
	for rows.Next() {
		update := models.NewPrincipalUpdate()
		err := rows.StructScan(update)
		if err != nil {
			return nil, err
		}
		err = validate.Struct(update)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)
	}
	return updates, nil
}

func InsertOrderInterestList(orderInterestList []*models.OrderInterest) error {
	sqlStr := `insert into OrderInterest (OrderID, Time, APY, InterestGain, TotalInterestGain) 
				values (:OrderID, :Time, :APY, :InterestGain, :TotalInterestGain)`
	_, err := SqlDB.NamedExec(sqlStr, orderInterestList)
	return err
}

func GetSortedOrderInterestListUntilDate(orderID string, until string) ([]*models.OrderInterest, error) {
	sqlStr := `select * from OrderInterest where OrderID = ? and Time <= ? order by Time asc`
	rows, err := SqlDB.Queryx(sqlStr, orderID, until)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	interestList := models.CreateOrderInterestList()
	for rows.Next() {
		interest := models.CreateOrderInterest()
		err := rows.StructScan(interest)
		if err != nil {
			return nil, err
		}
		err = validate.Struct(interest)
		if err != nil {
			return nil, err
		}
		interestList = append(interestList, interest)
	}
	return interestList, nil
}

func GetMostRecentOrderInterestUntilDate(orderID string, until string) (*models.OrderInterest, error) {
	interest := models.CreateOrderInterest()
	sqlStr := `select * from OrderInterest where OrderID = ? and Time <= ? order by Time desc`
	err := SqlDB.Get(interest, sqlStr, orderID, until)
	if err != nil {
		return nil, err
	}
	return interest, nil
}

func UpdateOrderAccumulatedInterest(orderID string, accumulatedInterest float64) error {
	sqlStr := "update Orders set AccumulatedInterest = ? where OrderID = ?"
	_, err := SqlDB.Exec(sqlStr, accumulatedInterest, orderID)
	return err
}

func UpdateOrderNewProductID(orderID string, newProductID string) error {
	sqlStr := `update Orders set ProductID = ? where OrderID = ?`
	_, err := SqlDB.Exec(sqlStr, newProductID, orderID)
	return err
}

/* 	DID           string
WalletAddress string
Type          string*/

func InsertMiningRole(role *models.MiningRole) error {
	sqlStr := `insert into MiningRole (DID, WalletAddress, Type ) 
				values (:DID, :WalletAddress, :Type)`
	_, err := SqlDB.NamedExec(sqlStr, role)
	return err
}

func GetMiningRole(DID string) (*models.MiningRole, error) {
	sqlStr := `select * from MiningRole where DID = ?`
	rows, err := SqlDB.Queryx(sqlStr, DID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var role models.MiningRole
		err := rows.StructScan(&role)
		if err != nil {
			return nil, err
		}
		return &role, nil
	}
	return nil, nil
}
