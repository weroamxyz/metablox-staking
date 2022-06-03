package dao

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"

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

func SubmitBuyinWithDate(tx *models.TXInfo) error {
	err := validate.Struct(tx)
	if err != nil {
		return err
	}
	dbTX, err := SqlDB.Beginx()
	if err != nil {
		return err
	}
	sqlStr := "update Orders set Type = 'Holding' where OrderID = ?"
	result, err := dbTX.Exec(sqlStr, tx.OrderID)
	if err != nil {
		dbTX.Rollback()
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		dbTX.Rollback()
		return err
	}
	if rows == 0 {
		dbTX.Rollback()
		return errval.ErrUpdateOrderStatus
	}

	sqlStr = "insert into TXInfo (OrderID, TXCurrencyType, TXType, TXHash, Principal, Interest, UserAddress, CreateDate, RedeemableTime) values (:OrderID, :TXCurrencyType, :TXType, :TXHash, :Principal, :Interest, :UserAddress, :CreateDate, :RedeemableTime)"
	_, err = dbTX.NamedExec(sqlStr, tx)
	if err != nil {
		dbTX.Rollback()
		return err
	}
	dbTX.Commit()
	return nil
}

func InsertPrincipalUpdateWithTime(productID string, totalPrincipal float64, time string) error {
	sqlStr := `insert into PrincipalUpdates (ProductID, Time, TotalPrincipal) values (?, ?, ?)`
	_, err := SqlDB.Exec(sqlStr, productID, time, totalPrincipal)
	return err
}

func BuyinTestOrderWithDate(order *models.Order, date string) (string, error) {
	id, err := CreateOrder(order)
	if err != nil {
		return "", err
	}

	// create corresponding txinfo
	txInfo := &models.TXInfo{
		OrderID:        strconv.Itoa(id),
		TXCurrencyType: "",
		TXType:         "BuyIn",
		TXHash:         nil,
		Principal:      order.Amount,
		Interest:       0,
		UserAddress:    "",
		CreateDate:     date,
		RedeemableTime: "2022-01-01 00:00:00",
	}
	err = SubmitBuyinWithDate(txInfo)
	if err != nil {
		return "", err
	}

	// record change in staking pool's total principal
	newPrincipal := models.NewPrincipalUpdate()
	oldPrincipal, err := GetLatestPrincipalUpdate(order.ProductID)
	if err == nil {
		newPrincipal.TotalPrincipal = oldPrincipal.TotalPrincipal + txInfo.Principal
	} else if err == sql.ErrNoRows {
		newPrincipal.TotalPrincipal = txInfo.Principal
	} else {
		return "", err
	}

	err = InsertPrincipalUpdateWithTime(order.ProductID, newPrincipal.TotalPrincipal, date)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(id), nil
}
