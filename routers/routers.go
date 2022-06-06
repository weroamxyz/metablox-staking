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
	r.GET("/staking/transactions/user/:did", controllers.GetTransactionsByUserDIDHandler)
	r.GET("/staking/interest/:id", controllers.GetOrderInterestHandler)
	r.POST("/staking/redeem/full/:id", controllers.RedeemOrderHandler)
	r.POST("/staking/redeem/interest/:id", controllers.RedeemInterestHandler)

	r.GET("/mining/minerlist", controllers.GetMinerListHandler)
	r.GET("/mining/miner", controllers.GetMinerByIDHandler)
	r.GET("/mining/exchangerate/:id", controllers.GetExchangeRateHandler)
	r.GET("/mining/rewardhistory/:did", controllers.GetRewardHistoryHandler)
	r.POST("/mining/exchange", controllers.ExchangeSeedHandler)
	r.POST("/mining/nonce", controllers.GetNonceHandler)
	r.POST("/mining/activate", controllers.ActivateExchangeHandler)
	r.POST("/mining/newexchange", controllers.NewSeedExchangeHandler)

	r.Run(":8889")
}
