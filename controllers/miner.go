package controllers

import (
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
)

func GetClosestMiner(latitude, longitude string) (*models.MinerInfo, error) {
	minerList, err := GetAllMiners()
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

func GetMinerList(c *gin.Context) ([]*models.MinerInfo, error) {

	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	if latitude == "" || longitude == "" {
		minerList, err := GetAllMiners()
		if err != nil {
			return nil, err
		}
		return minerList, nil
	}

	minerList, err := dao.GetAllVirtualMinerInfo() //return all virtual miners along with the closest one
	if err != nil {
		return nil, err
	}

	closestMiner, err := GetClosestMiner(latitude, longitude)
	if err != nil {
		return nil, err
	}

	minerList = append(minerList, closestMiner)

	return minerList, nil
}

func GetAllMiners() ([]*models.MinerInfo, error) {
	minerList, err := dao.GetAllMinerInfo()
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

	minerID := c.Query("minerid")

	miner, err := dao.GetMinerInfoByID(minerID)
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
