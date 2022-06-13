package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/metabloxStaking/models"
)

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

func GetAllVirtualMinerInfo() ([]*models.MinerInfo, error) {
	var miners []*models.MinerInfo
	sqlStr := "select * from MinerInfo where IsVirtual = 1"
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
	sqlStr := "insert into SeedExchangeHistory (UserDID, TargetDID, Challenge, ExchangeRate, Amount) values (:UserDID, :TargetDID, :Challenge, :ExchangeRate, :Amount)"
	_, err := SqlDB.NamedExec(sqlStr, exchange)
	return err
}

func GetExchangeRate(id string) (float64, error) {
	var rate float64
	sqlStr := "select ExchangeRate from ExchangeRate where ID = ?"
	err := SqlDB.Get(&rate, sqlStr, id)
	return rate, err
}

func CheckIfDIDIsValidator(did string) (bool, error) {
	var count int
	sqlStr := "select count(*) from WifiAccessInfo where ID = ? and Type = 'Validator'"
	err := SqlDB.Get(&count, sqlStr, did)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}

func CheckIfDIDIsMiner(did string) (bool, error) {
	var count int
	sqlStr := "select count(*) from MiningLicenseInfo where ID = ?"
	err := SqlDB.Get(&count, sqlStr, did)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}
