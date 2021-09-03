package shippingVendorRepository

import (
	"ecommerce/models"
	"gorm.io/gorm"
)

func FindPrice(db *gorm.DB, fromCity string, toCity string, weight float32) (float32, int64) {
	var ShippingVendor models.ShippingVendors
	query := db.Model(&models.ShippingVendors{}).Where("from_city = ?", fromCity).Where("to_city = ?", toCity).Find(&ShippingVendor)
	if query.RowsAffected > 0 {
		return ShippingVendor.PricePerKg * weight / 1000, query.RowsAffected
	} else {
		return 0, 0
	}
}
