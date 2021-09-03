package categoryRepository

import (
	"ecommerce/models"
	"gorm.io/gorm"
)

func FindAll(db *gorm.DB) ([]models.Categories, int64, error) {
	var Categories []models.Categories
	query := db.Find(&Categories)
	return Categories, query.RowsAffected, query.Error
}
