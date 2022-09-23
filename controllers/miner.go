package controllers

import (
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
)

const acceptableDistance = 30 //arbitrary value; can be increased to increase search range for miners, or decreased to do the opposite

//returns all non-virtual miners which are within a certain distance of the provided coordinates
func GetNearbyMiners(latitude, longitude string) ([]*models.MinerInfo, error) {
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

	var nearbyMiners []*models.MinerInfo

	for _, miner := range minerList {
		if miner.Longitude == nil || miner.Latitude == nil {
			continue
		}
		longDistance := floatLong - *miner.Longitude
		latDistance := floatLat - *miner.Latitude
		totalDistance := math.Sqrt(math.Pow(longDistance, 2) + math.Pow(latDistance, 2)) //use pythagorean theorem to calculate distance
		if totalDistance < acceptableDistance {
			nearbyMiners = append(nearbyMiners, miner)
		}
	}
	return nearbyMiners, nil
}

//returns all virtual miners as well as any within a certain distance of the provided coordinates
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

	minerList, err := dao.GetAllVirtualMinerInfo()
	if err != nil {
		return nil, err
	}

	nearbyMiners, err := GetNearbyMiners(latitude, longitude)
	if err != nil {
		return nil, err
	}

	minerList = append(minerList, nearbyMiners...)

	return minerList, nil
}

//return all miners
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

//return the miner that has the specified ID
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

//return the miner that has the specified BSSID
func GetMinerByBSSID(c *gin.Context) (*models.MinerInfo, error) {

	bssid := c.Param("BSSID")
	println(bssid)

	miner, err := dao.GetMinerInfoByBSSID(bssid)
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
