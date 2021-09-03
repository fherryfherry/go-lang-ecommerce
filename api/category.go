package api

import (
	"ecommerce/database"
	"ecommerce/repositories/categoryRepository"
	"github.com/gin-gonic/gin"
)

func CategoryList(c *gin.Context) {
	db := database.Connect(c)

	Categories, _, err := categoryRepository.FindAll(db)
	if err != nil {
		c.JSON(500, gin.H{"status":0,"message":err.Error()})
		return
	}

	c.JSON(200, gin.H{"status":1,"message":"ok","data": Categories})
}
