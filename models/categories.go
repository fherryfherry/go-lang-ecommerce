package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Name	string	`json:"name" gorm:"size:255"`
}