package userRepository

import (
	"ecommerce/models"
	"gorm.io/gorm"
)

func FindByEmail(db *gorm.DB, email string) (models.Users, int64) {
	var User models.Users
	query := db.First(&User, "email = ?", email)
	return User, query.RowsAffected
}
