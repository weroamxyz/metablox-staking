package dao

import (
	"fmt"
	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

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
	product := models.NewStakingProduct()

	sqlStr := "select * from StakingProducts where ID = ?"
	err := SqlDB.Get(product, sqlStr, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func GetAllProductInfo() ([]*models.StakingProduct, error) {
	var products []*models.StakingProduct
	sqlStr := "select * from StakingProducts"
	rows, err := SqlDB.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.NewStakingProduct()
		err = rows.StructScan(product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

func GetStakingRecords(did string) ([]*models.StakingRecord, error) {
	var records []*models.StakingRecord
	sqlStr := "select Orders.OrderID, Orders.Term, Orders.Type, TXInfo.CreateDate from Orders join TXInfo on TXInfo.OrderID = Orders.OrderID where Orders.UserDID = ? and TXInfo.TXType = 'BuyIn'"
	rows, err := SqlDB.Queryx(sqlStr, did)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		record := models.NewStakingRecord()
		err = rows.StructScan(record)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func PrepareGetInterestByOrderID() (*sqlx.Stmt, error) {
	sqlStr := "select TotalInterestGain from OrderInterest where OrderID = ? order by ID desc limit 1"
	stmt, err := SqlDB.Preparex(sqlStr)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func ExecuteGetInterestStmt(id int, stmt *sqlx.Stmt) (float32, error) {
	var interest float32
	err := stmt.Get(&interest, id)
	if err != nil {
		return 0, err
	}
	return interest, nil
}

func GetTransactionsByOrderID(orderID string) ([]*models.TXInfo, error) {
	var transactions []*models.TXInfo
	sqlStr := "select * from TXInfo where OrderID = ?"
	rows, err := SqlDB.Queryx(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tx := models.NewTXInfo()
		err = rows.StructScan(tx)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, err
}

func RedeemOrderByOrderID(orderID string) error {
	sqlStr := "update Orders set RedeemStatus = 1 where OrderID = ?"
	_, err := SqlDB.Exec(sqlStr, orderID)
	if err != nil {
		return err
	}
	return nil
}

func RedeemInterestByOrderID(orderID string) error {
	sqlStr := "update OrderInterest set TotalInterestGain = 0 where OrderID = ? order by ID desc limit 1"
	_, err := SqlDB.Exec(sqlStr, orderID)
	if err != nil {
		return err
	}
	return nil
}

func GetOrderByID(orderID string) (*models.Order, error) {
	order := models.NewOrder()
	sqlStr := "select * from Orders where OrderID = ?"
	err := SqlDB.Get(order, sqlStr, orderID)
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

func GetOrderBuyInPrincipal(orderID string) (float32, error) {
	var buyInAmount float32
	sqlStr := "select Principal from TXInfo where OrderID = ? and TXType = 'BuyIn'"
	err := SqlDB.Get(&buyInAmount, sqlStr, orderID)
	if err != nil {
		return 0.0, err
	}
	return buyInAmount, nil
}

func CheckIfOrderMeetsMinimumInterest(orderID string) (bool, error) {
	var minInterest int
	sqlStr := "select StakingProducts.MinRedeemValue from StakingProducts join Orders on StakingProducts.ID = Orders.ProductID where Orders.OrderID = ?"
	err := SqlDB.Get(&minInterest, sqlStr, orderID)
	if err != nil {
		return false, err
	}

	var currentInterest float32
	sqlStr = "select TotalInterestGain from OrderInterest where OrderID = ? order by ID desc limit 1"
	err = SqlDB.Get(&currentInterest, sqlStr, orderID)
	if err != nil {
		return false, err
	}

	if int(currentInterest) < minInterest { //todo: include proper conversion since I think these values are in different units
		return false, nil
	}

	return true, nil
}

func UploadTransaction(tx *models.TXInfo) error {
	sqlStr := "insert into TXInfo (UserDID, OrderID, TXCurrencyTYPE, TXType, TXHash, Amount, UserAddress) values (:UserDID, :OrderID, :TXCurrencyTYPE, :TXType, :TXHash, :Amount, :UserAddress)"
	_, err := SqlDB.NamedExec(sqlStr, tx)
	if err != nil {
		return err
	}
	return nil
}

func InsertPrincipalUpdate(productID int, totalPrincipal float32) error {
	sqlStr := `insert into PrincipalUpdates (ProductID, TotalPrincipal) values (?, ?)`
	_, err := SqlDB.Exec(sqlStr, productID, totalPrincipal)
	return err
}

func GetPrincipalUpdates(productID int) ([]*models.PrincipalUpdate, error) {
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
	sqlStr := `insert into OrderInterest (OrderID, Time, APY, InterestGain) values (:OrderID, :Time, :APY, :InterestGain)`
	_, err := SqlDB.NamedExec(sqlStr, orderInterestList)
	return err
}

func GetSortedOrderInterestList(orderID int) ([]*models.OrderInterest, error) {
	sqlStr := `select * from OrderInterest where OrderID = ? order by Time asc`
	rows, err := SqlDB.Queryx(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	interestList := models.NewOrderInterestList()
	for rows.Next() {
		interest := models.NewOrderInterest()
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
