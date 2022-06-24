package dao

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
)

var sqlUser = "admin"
var sqlPassword = "password"
var sqlHost = "127.0.0.1"
var sqlPort = "3306"
var sqlDBName = "metabloxstakingtest"

const SQLCreateScript = "../dao/create_test_tables.sql"
const SQLInsertScript = "../dao/insert_test_data.sql"
const SQLDeleteScript = "../dao/delete_test_tables.sql"

func InitTestDB() error {
	validate = validator.New()
	if err := connectTestDB(); err != nil {
		return err
	}
	if err := executeSQLScript(SQLCreateScript); err != nil {
		return err
	}
	if err := executeSQLScript(SQLInsertScript); err != nil {
		return err
	}
	return nil
}

func CleanupTestDB() {
	defer SqlDB.Close()
	err := executeSQLScript(SQLDeleteScript)
	if err != nil {
		logger.Warn(err.Error())
	}
}

func connectTestDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		sqlUser,
		sqlPassword,
		sqlHost,
		sqlPort,
		sqlDBName,
	)
	SqlDB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

func executeSQLScript(path string) error {
	scriptBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = SqlDB.Exec(string(scriptBytes))
	if err != nil {
		return err
	}
	return nil
}

func insertPrincipalUpdateWithTime(productID string, totalPrincipal string, time string) error {
	sqlStr := `insert into PrincipalUpdates (ProductID, Time, TotalPrincipal) values (?, ?, ?)`
	_, err := SqlDB.Exec(sqlStr, productID, time, totalPrincipal)
	return err
}

func BuyinTestOrderWithDate(order *models.Order, date string) (string, error) {
	id, err := CreateOrder(order)
	if err != nil {
		return "", err
	}
	orderID := strconv.Itoa(id)

	// create corresponding txinfo
	txInfo := &models.TXInfo{
		OrderID:        strconv.Itoa(id),
		TXType:         models.TxTypeBuyIn,
		TXHash:         "",
		Principal:      order.Amount,
		CreateDate:     date,
		RedeemableTime: date,
	}
	err = SubmitBuyin(txInfo)
	if err != nil {
		return "", err
	}

	err = SetTXInfoDate(orderID, models.TxTypeBuyIn, date)
	if err != nil {
		return "", err
	}

	// record change in staking pool's total principal
	newPrincipal := models.NewPrincipalUpdate()
	oldPrincipal, err := GetLatestPrincipalUpdate(order.ProductID)
	if err == nil {
		newPrincipal.TotalPrincipal = oldPrincipal.TotalPrincipal.Add(txInfo.Principal)
	} else if err == sql.ErrNoRows {
		newPrincipal.TotalPrincipal = txInfo.Principal
	} else {
		return "", err
	}

	err = insertPrincipalUpdateWithTime(order.ProductID, newPrincipal.TotalPrincipal.String(), date)
	if err != nil {
		return "", err
	}
	return orderID, nil
}

func RedeemTestOrderWithDate(orderID string, date string) error {
	interestInfo, err := GetInterestInfoByOrderID(orderID)
	if err != nil {
		return err
	}

	order, err := GetOrderByID(orderID)
	if err != nil {
		return err
	}

	// create corresponding txinfo
	txInfo := &models.TXInfo{
		OrderID:        orderID,
		TXType:         models.TxTypeOrderClosure,
		TXHash:         "",
		Principal:      order.Amount,
		CreateDate:     date,
		RedeemableTime: date,
	}

	err = RedeemOrder(txInfo, interestInfo.AccumulatedInterest.String())
	if err != nil {
		return err
	}

	err = SetTXInfoDate(orderID, models.TxTypeOrderClosure, date)
	if err != nil {
		return err
	}

	// record change in staking pool's total principal
	newPrincipal := models.NewPrincipalUpdate()
	oldPrincipal, err := GetLatestPrincipalUpdate(order.ProductID)
	if err != nil {
		return err
	}
	newPrincipal.TotalPrincipal = oldPrincipal.TotalPrincipal.Sub(order.Amount)

	err = insertPrincipalUpdateWithTime(order.ProductID, newPrincipal.TotalPrincipal.String(), date)
	if err != nil {
		return err
	}
	return nil
}

func SetTXInfoDate(orderID string, txType string, date string) error {
	sqlStr := `update TXInfo set CreateDate = ? where OrderID = ? and TXType = ?`
	_, err := SqlDB.Exec(sqlStr, date, orderID, txType)
	return err
}
