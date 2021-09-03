package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	SKU 			string 		`json:"sku" gorm:"size:255"`
	Name 			string 		`json:"name" gorm:"size:255"`
	Description 	string 		`json:"description" gorm:"size:1000"`
	Price			float32		`json:"price" gorm:"size:255"`
	CategoriesID	int			`json:"category_id"`
	Categories		Categories	`gorm:"constraint:OnDelete:SET NULL"`
}