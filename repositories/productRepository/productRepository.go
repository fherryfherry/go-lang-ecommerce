package productRepository

import (
	"ecommerce/models"
	"gorm.io/gorm"
	"strconv"
)

func FindAll(db *gorm.DB, orderBy string, orderDir string, limit int, offset string) ([]models.Products, int64) {
	var User []models.Products
	offsetInt, _ := strconv.Atoi(offset)
	query := db.Model(&models.Products{}).Where("status = 'Active'").Limit(limit).Offset(offsetInt).Order(orderBy + " " + orderDir).Find(&User)
	return User, query.RowsAffected
}

func FindById(db *gorm.DB, ID int) (models.Products, int64) {
	var Product models.Products
	query := db.Model(&models.Products{}).Where("id = ?", ID).Where("status = 'Active'").First(&Product)
	return Product, query.RowsAffected
}
