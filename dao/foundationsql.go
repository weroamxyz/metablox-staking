package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
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

func GetMinerInfoByBSSID(bssid string) (*models.MinerInfo, error) {
	miner := models.CreateMinerInfo()
	sqlStr := "select * from MinerInfo where BSSID = ?"
	err := SqlDB.Get(miner, sqlStr, bssid)
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

func CheckIfExchangeExists(keys *models.SeedHistoryKeys) error {
	var count int
	sqlStr := "select count(*) from SeedExchangeHistory where UserDID = ? and TargetDID = ? and Challenge = ?"
	err := SqlDB.Get(&count, sqlStr, keys.DID, keys.Target, keys.Challenge)
	if err != nil {
		return err
	}
	if count != 0 {
		return errval.ErrSeedAlreadyExchanged
	}
	return nil
}

func GetExchangeRate(id string) (decimal.Decimal, error) {
	var stringRate string
	sqlStr := "select ExchangeRate from ExchangeRate where ID = ?"
	err := SqlDB.Get(&stringRate, sqlStr, id)
	if err != nil {
		return decimal.Decimal{}, err
	}
	rate, err := decimal.NewFromString(stringRate)
	if err != nil {
		logger.Warn("exec GetExchangeRate function error:", err)
		return decimal.Decimal{}, errval.ErrExchangeRateNotNumber
	}

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

func GetWifiAccessInfo(id string) (*models.WifiAccessInfo, error) {
	miner := models.CreateWifiAccessInfo()
	sqlStr := "select * from WifiAccessInfo where ID = ?"
	err := SqlDB.Get(miner, sqlStr, id)
	if err != nil {
		return nil, err
	}

	return miner, nil
}
