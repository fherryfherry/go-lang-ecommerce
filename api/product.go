package api

import (
	"ecommerce/database"
	"ecommerce/models"
	"ecommerce/repositories/productRepository"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ProductList(c *gin.Context) {
	db := database.Connect(c)

	var result []models.Products

	CategoryID := c.Query("categories_id")
	if len(CategoryID) > 0 {
		CategoryIDInt, _ := strconv.Atoi(CategoryID)
		result, _ = productRepository.FindAllByCategory(db, CategoryIDInt, "id", "desc", 10, c.DefaultQuery("offset", "0"))
	} else {
		result, _ = productRepository.FindAll(db, "id", "desc", 10, c.DefaultQuery("offset", "0"))
	}

	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": result})
}

func ProductDetail(c *gin.Context) {
	db := database.Connect(c)
	ID, _ := strconv.Atoi(c.Param("id"))
	row, _ := productRepository.FindById(db, ID)
	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": row})
}
