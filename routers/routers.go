package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/controllers"
	"github.com/metabloxStaking/log"
	"github.com/metabloxStaking/middleware"
)

func Setup() {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(log.GetLogWriter()), gin.RecoveryWithWriter(log.GetLogWriter()))

	didParamGroup := r.Group("/")
	didQueryGroup := r.Group("/")

	r.GET("/product/search/:id", controllers.GetProductInfoByIDHandler)

	r.GET("/product/all", controllers.GetAllProductInfoHandler)
	r.POST("/order/create", controllers.CreateOrderHandler)
	r.POST("/order/confirm", controllers.SubmitBuyinHandler)

	r.GET("/staking/transactions/order/:id", controllers.GetTransactionsByOrderIDHandler)
	r.GET("/staking/interest/:id", controllers.GetOrderInterestHandler)
	r.POST("/staking/redeem/full/:id", controllers.RedeemOrderHandler)
	r.POST("/staking/redeem/interest/:id", controllers.RedeemInterestHandler)

	r.GET("/mining/exchangerate/:id", controllers.GetExchangeRateHandler)
	//r.POST("/mining/exchange", controllers.ExchangeSeedHandler)
	r.POST("/mining/nonce", controllers.GetNonceHandler)
	r.POST("/mining/activate", controllers.ActivateExchangeHandler)
	r.POST("/mining/newexchange", controllers.NewSeedExchangeHandler)

	r.GET("/wifi-profile/ios", controllers.GetIOSprofileHandler)
	r.GET("/wifi/:account", controllers.GetWifiAccessInfoHandler)
	r.GET("/miners", controllers.GetMinerListHandler)
	r.GET("/miner/:BSSID", controllers.GetMinerByBSSIDHandler)

	didParamGroup.Use(middleware.DIDParamMiddleware())
	{
		didParamGroup.GET("/staking/orders/:did", controllers.GetStakingRecordsHandler)
		didParamGroup.GET("/staking/transactions/user/:did", controllers.GetTransactionsByUserDIDHandler)
		didParamGroup.GET("/mining/rewardhistory/:did", controllers.GetRewardHistoryHandler)
	}

	didQueryGroup.Use(middleware.DIDQueryMiddleware())
	{
		didQueryGroup.GET("/mining/minerlist", controllers.GetMinerListHandler)
		didQueryGroup.GET("/mining/miner", controllers.GetMinerByIDHandler)
	}

	r.Run(":8886")
}
