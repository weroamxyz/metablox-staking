package foundationdao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var SqlDB *sqlx.DB

func InitSql() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString("foundationmysql.user"),
		viper.GetString("foundationmysql.password"),
		viper.GetString("foundationmysql.host"),
		viper.GetString("foundationmysql.port"),
		viper.GetString("foundationmysql.dbname"),
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

func GetAllMinerInfo() ([]*models.MinerInfo, error) {
	var miners []*models.MinerInfo
	sqlStr := "select * from MinerInfo"
	rows, err := SqlDB.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		miner := models.CreateMinerInfo()
		err = rows.StructScan(miner)
		if err != nil {
			return nil, err
		}
		miners = append(miners, miner)
	}
	return miners, nil
}

func GetMinerInfoByID(id string) (*models.MinerInfo, error) {
	miner := models.CreateMinerInfo()
	sqlStr := "select * from MinerInfo where ID = ?"
	err := SqlDB.Get(miner, sqlStr, id)
	if err != nil {
		return nil, err
	}

	return miner, nil
}

func GetSeedHistory(did string) ([]*models.SeedExchange, error) {
	var exchangeList []*models.SeedExchange
	sqlStr := "select * from SeedExchangeHistory where UserDID = ?"
	rows, err := SqlDB.Queryx(sqlStr, did)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		exchange := models.CreateSeedExchange()
		err = rows.StructScan(exchange)
		if err != nil {
			return nil, err
		}
		exchangeList = append(exchangeList, exchange)
	}
	return exchangeList, nil
}

func UploadSeedExchange(exchange *models.SeedExchange) error {
	sqlStr := "insert into SeedExchangeHistory (VcID, UserDID, ExchangeRate, Amount) values (:VcID, :UserDID, :ExchangeRate, :Amount)"
	_, err := SqlDB.NamedExec(sqlStr, exchange)
	return err
}

func GetExchangeRate(id string) (float64, error) {
	var rate float64
	sqlStr := "select ExchangeRate from ExchangeRate where ID = ?"
	err := SqlDB.Get(&rate, sqlStr, id)
	return rate, err
}
