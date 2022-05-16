package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/controllers"
)

func Setup() {
	r := gin.New()

	r.GET("/product/search/:id", controllers.GetProductInfoByIDHandler)

	r.GET("/product/all", controllers.GetAllProductInfoHandler)
	r.POST("/purchase/:id", controllers.PurchaseProductByIDHandler)

	r.GET("/records/:did", controllers.GetStakingRecordsHandler)
	r.GET("/transactions/:id", controllers.GetTransactionsByOrderIDHandler)
	r.POST("/redeem/full", controllers.RedeemOrderHandler)
	r.POST("/redeem/interest", controllers.RedeemInterestHandler)
	r.Run(":8889")
}
