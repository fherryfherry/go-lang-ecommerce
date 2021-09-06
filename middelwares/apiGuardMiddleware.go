package middelwares

import (
	"ecommerce/database"
	"ecommerce/helpers"
	"ecommerce/repositories/tokenRepository"
	"github.com/gin-gonic/gin"
)

func ApiGuardMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := helpers.GetAuthorizationToken(c)
		if accessToken != "" {
			db := database.Connect(c)
			_, found := tokenRepository.FindByAccessToken(db, accessToken)
			if found == 0 {
				c.JSON(401, gin.H{"status": 0, "message": "Access token expired!"})
				c.Abort()
				return
			}

			c.Next()
		} else {
			c.JSON(401, gin.H{"status": 0, "message": "Invalid access!"})
			c.Abort()
			return
		}
	}
}
