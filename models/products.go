package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	SKU          string     `json:"sku" gorm:"size:255"`
	Name         string     `json:"name" gorm:"size:255"`
	Description  string     `json:"description" gorm:"size:1000"`
	Price        float32    `json:"price" gorm:"size:255"`
	Stock        int        `json:"stock" gorm:"size:11"`
	Status       string     `json:"status" gorm:"size:25"`
	CategoriesID int        `json:"categories_id"`
	Categories   Categories `gorm:"constraint:OnDelete:SET NULL"`
	BrandsID     int        `gorm:"brands_id"`
	Brands       Brands     `gorm:"constraint:OnDelete:SET NULL"`
}
