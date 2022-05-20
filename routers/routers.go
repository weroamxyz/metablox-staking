package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/controllers"
)

func Setup() {
	r := gin.New()

	r.GET("/product/search/:id", controllers.GetProductInfoByIDHandler)

	r.GET("/product/all", controllers.GetAllProductInfoHandler)
	r.POST("/order/create", controllers.CreateOrderHandler)
	r.POST("/order/confirm", controllers.SubmitBuyinHandler)

	r.GET("/staking/orders/:did", controllers.GetStakingRecordsHandler)
	r.GET("/staking/transactions/order/:id", controllers.GetTransactionsByOrderIDHandler)
	r.GET("/staking/transactions/user/:id", controllers.GetTransactionsByUserDIDHandler)
	r.GET("/staking/interest/:id", controllers.GetOrderInterestHandler)
	r.POST("/staking/redeem/full/:id", controllers.RedeemOrderHandler)
	r.POST("/staking/redeem/interest/:id", controllers.RedeemInterestHandler)
	r.Run(":8889")
}
