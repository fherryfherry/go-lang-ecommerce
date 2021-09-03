package models

import "gorm.io/gorm"

type ShippingVendors struct {
	gorm.Model
	Name       string  `json:"name" gorm:"size:255"`
	FromCity   string  `json:"from_city" gorm:"size:255"`
	ToCity     string  `json:"to_city" gorm:"size:255"`
	PricePerKg float32 `json:"price_per_gram"`
}
