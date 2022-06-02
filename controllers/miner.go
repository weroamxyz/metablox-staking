package controllers

import (
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/foundationdao"
	"github.com/metabloxStaking/models"
)

func GetClosestMiner(latitude, longitude string) (*models.MinerInfo, error) {
	minerList, err := GetMinerList()
	if err != nil {
		return nil, err
	}

	floatLat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return nil, err
	}

	floatLong, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil, err
	}

	closestMiner := models.CreateMinerInfo()
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
	return closestMiner, nil
}

func GetMinerList() ([]*models.MinerInfo, error) {
	minerList, err := foundationdao.GetAllMinerInfo()
	if err != nil {
		return nil, err
	}

	for _, miner := range minerList {
		createDate, err := time.Parse("2006-01-02 15:04:05", miner.CreateTime)
		if err != nil {
			return nil, err
		}
		miner.CreateTime = strconv.FormatFloat(float64(createDate.UnixNano())/float64(time.Second), 'f', 3, 64)
	}

	return minerList, nil
}

func GetMinerByID(c *gin.Context) (*models.MinerInfo, error) {
	did := c.Query("did")
	err := validateDID(did)
	if err != nil {
		return nil, err
	}

	minerID := c.Query("minerid")

	miner, err := foundationdao.GetMinerInfoByID(minerID)
	if err != nil {
		return nil, err
	}

	createDate, err := time.Parse("2006-01-02 15:04:05", miner.CreateTime)
	if err != nil {
		return nil, err
	}
	miner.CreateTime = strconv.FormatFloat(float64(createDate.UnixNano())/float64(time.Second), 'f', 3, 64)

	return miner, nil
}
