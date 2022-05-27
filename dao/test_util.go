package dao

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

var sqlUser = "admin"
var sqlPassword = "password"
var sqlHost = "127.0.0.1"
var sqlPort = "3306"
var sqlDBName = "boundarypaytest"

const SQLCreateScript = "../dao/create_test_tables.sql"
const SQLDeleteScript = "../dao/delete_test_tables.sql"

func InitTestDB() error {
	if err := connectTestDB(); err != nil {
		return err
	}
	if err := executeTableCreateScript(); err != nil {
		return err
	}
	return nil
}

func CleanupTestDB() {
	defer SqlDB.Close()
	err := executeTableDeleteScript()
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

func executeTableCreateScript() error {
	scriptBytes, err := ioutil.ReadFile(SQLCreateScript)
	if err != nil {
		return err
	}
	_, err = SqlDB.Exec(string(scriptBytes))
	if err != nil {
		return err
	}
	return nil
}

func executeTableDeleteScript() error {
	scriptBytes, err := ioutil.ReadFile(SQLDeleteScript)
	if err != nil {
		return err
	}
	_, err = SqlDB.Exec(string(scriptBytes))
	if err != nil {
		return err
	}
	return nil
}
