package middelwares

import (
	"ecommerce/helpers"
	"github.com/gin-gonic/gin"
)

func ApiGuardMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := helpers.GetAuthorizationToken(c)
		if accessToken != "" {
			c.Next()
		} else {
			c.JSON(401, gin.H{"status": 0, "message": "Invalid access!"})
			c.Abort()
			return
		}
	}
}
