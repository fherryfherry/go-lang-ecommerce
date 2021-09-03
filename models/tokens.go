package models

import (
	"gorm.io/gorm"
	"time"
)

type Tokens struct {
	gorm.Model
	AccessToken		string		`json:"access_token" gorm:"size:255"`
	RefreshToken	string		`json:"refresh_token" gorm:"size:255"`
	IpAddress		string		`json:"ip_address" gorm:"size:255"`
	UserAgent		string		`json:"user_agent" gorm:"size:255"`
	ExpiredAt		time.Time	`json:"expired_at" gorm:"size:255"`
	UsersID			int 		`json:"users_id" gorm:"default:NULL"`
	Users			Users		`gorm:"constraint:OnDelete:CASCADE"`
}
