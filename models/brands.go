package models

import "gorm.io/gorm"

type Brands struct {
	gorm.Model
	Name string `json:"name" gorm:"size:255"`
}
