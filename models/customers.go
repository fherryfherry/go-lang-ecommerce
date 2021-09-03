package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name       string `json:"name" gorm:"size:255"`
	Email      string `json:"email" gorm:"size:255"`
	Phone      string `json:"phone" gorm:"size:255"`
	Address    string `json:"address" gorm:"size:255"`
	City       string `json:"city" gorm:"size:255"`
	District   string `json:"district" gorm:"size:255"`
	PostalCode string `json:"postal_code" gorm:"size:255"`
}
