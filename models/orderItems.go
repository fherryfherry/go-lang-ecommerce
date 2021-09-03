package models

import (
	"gorm.io/gorm"
)

type OrderItems struct {
	gorm.Model
	OrdersID   uint     `json:"orders_id"`
	Orders     Orders   `gorm:"constraint:OnDelete:CASCADE"`
	ProductsID uint     `json:"products_id"`
	Products   Products `gorm:"constraint:OnDelete:CASCADE"`
	Price      float32  `json:"price" gorm:"default:0"`
	Qty        float32  `json:"qty" gorm:"default:0"`
	SubTotal   float32  `json:"sub_total" gorm:"default:0"`
}
