package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
)

func GetWifiByAccount(c *gin.Context) (*models.WifiAccessInfo, error) {
	id := c.Param("account")
	println(id)
	wifi, err := dao.GetWifiAccessInfo(id)

	if err != nil {
		return nil, err
	}

	return wifi, nil

}
