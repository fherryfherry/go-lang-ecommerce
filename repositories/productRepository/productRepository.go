package productRepository

import (
	"ecommerce/models"
	"gorm.io/gorm"
	"strconv"
)

func FindAll(db *gorm.DB, orderBy string, orderDir string, limit int, offset string) ([]models.Products, int64) {
	var User []models.Products
	offsetInt, _ := strconv.Atoi(offset)
	query := db.Model(&models.Products{})
	query = query.Where("status = 'Active'")
	query = query.Limit(limit)
	query = query.Offset(offsetInt).Order(orderBy + " " + orderDir).Find(&User)
	return User, query.RowsAffected
}

func FindAllByCategory(db *gorm.DB, categoryID int, orderBy string, orderDir string, limit int, offset string) ([]models.Products, int64) {
	var User []models.Products
	offsetInt, _ := strconv.Atoi(offset)
	query := db.Model(&models.Products{})
	query = query.Where("status = 'Active'")
	query = query.Where("categories_id = ?", categoryID)
	query = query.Limit(limit)
	query = query.Offset(offsetInt).Order(orderBy + " " + orderDir).Find(&User)
	return User, query.RowsAffected
}

func FindById(db *gorm.DB, ID uint) (models.Products, int64) {
	var Product models.Products
	query := db.Model(&models.Products{}).Where("id = ?", ID).Where("status = 'Active'").First(&Product)
	return Product, query.RowsAffected
}
