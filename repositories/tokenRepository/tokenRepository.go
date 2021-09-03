package tokenRepository

import (
	"ecommerce/models"
	"gorm.io/gorm"
)

func FindByAccessToken(db *gorm.DB, accessToken string) (models.Tokens, int64) {
	var token models.Tokens
	query := db.First(&token, "access_token = ?", accessToken)
	return token, query.RowsAffected
}

func FindByRefreshToken(db *gorm.DB, refreshToken string) (models.Tokens, int64) {
	var token models.Tokens
	query := db.First(&token, "refresh_token = ?", refreshToken)
	return token, query.RowsAffected
}
