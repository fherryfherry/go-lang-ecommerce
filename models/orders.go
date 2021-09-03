package models

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderNo         string `json:"order_no" gorm:"size:255"`
	CustomersID     int    `json:"customer_id"`
	Customers       Customers
	ShippingVendor  string  `json:"shipping_vendor" gorm:"size:255;default:NULL"`
	ShippingPackage string  `json:"shipping_package" gorm:"size:255;default:NULL"`
	ShippingPrice   float32 `json:"shipping_price" gorm:"default:0"`
	GrandTotal      float32 `json:"grand_total" gorm:"default:0"`
	PaymentStatus   float32 `json:"payment_status" gorm:"size:25"`
	DeliveryStatus  float32 `json:"delivery_status" gorm:"size:25"`
	OrderStatus     float32 `json:"order_status" gorm:"size:25"`
}
