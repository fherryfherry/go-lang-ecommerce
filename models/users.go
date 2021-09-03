package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name		string	`json:"name" gorm:"size:255"`
	Photo		string	`json:"photo" gorm:"size:255"`
	Email		string	`json:"email" gorm:"size:255"`
	Password	string	`json:"password" gorm:"size:255"`
}

func (s Users) getName() string {
	return s.Name
}

func (s Users) getPhoto() string {
	return s.Photo
}