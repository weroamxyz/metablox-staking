package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/controllers"
)

func DIDParamMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		did := c.Param("did")
		if controllers.ValidateDID(did) != nil {
			controllers.ResponseError(c, controllers.CodeInvalidAuth)
			c.Abort()
			return
		}
		c.Next()
	}
}

func DIDQueryMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		did := c.Query("did")
		if controllers.ValidateDID(did) != nil {
			controllers.ResponseError(c, controllers.CodeInvalidAuth)
			c.Abort()
			return
		}
		c.Next()
	}
}
