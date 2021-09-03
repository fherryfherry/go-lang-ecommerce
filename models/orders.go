package models

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderNo             string  `json:"order_no" gorm:"size:255"`
	CustomersID         int     `json:"users_id"`
	Customers           Users   `gorm:"constraint:OnDelete:CASCADE"`
	RecipientName       string  `json:"recipient_name" gorm:"size:255"`
	RecipientAddress    string  `json:"recipient_address" gorm:"size:300"`
	RecipientCity       string  `json:"recipient_city" gorm:"size:255"`
	RecipientPostalCode string  `json:"recipient_postal_code" gorm:"size:25"`
	ShippingVendor      string  `json:"shipping_vendor" gorm:"size:255;default:NULL"`
	ShippingPackage     string  `json:"shipping_package" gorm:"size:255;default:NULL"`
	ShippingPrice       float32 `json:"shipping_price" gorm:"default:0"`
	TotalWeight         float32 `json:"total_weight"`
	GrandTotal          float32 `json:"grand_total" gorm:"default:0"`
	PaymentStatus       string  `json:"payment_status" gorm:"size:25"`
	DeliveryStatus      string  `json:"delivery_status" gorm:"size:25"`
	OrderStatus         string  `json:"order_status" gorm:"size:25"`
}
