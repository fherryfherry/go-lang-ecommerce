package helpers

import "github.com/gin-gonic/gin"

func ResponseErrorIfAny(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, gin.H{"status":0,"message": err.Error()})
		return
	}
}
