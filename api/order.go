package api

import (
	"ecommerce/database"
	"ecommerce/helpers"
	"ecommerce/repositories/tokenRepository"
	"ecommerce/services/orderService"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type OrderItemParameters struct {
	ProductsID uint    `json:"products_id" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	Qty        float32 `json:"qty" binding:"required"`
}

type OrderParameters struct {
	ShippingVendor      string                `json:"shipping_vendor" binding:"required"`
	ShippingPackage     string                `json:"shipping_package" binding:"required"`
	RecipientName       string                `json:"recipient_name" binding:"required"`
	RecipientAddress    string                `json:"recipient_address" binding:"required"`
	RecipientCity       string                `json:"recipient_city" binding:"required"`
	RecipientPostalCode string                `json:"recipient_postal_code" binding:"required"`
	OrderItems          []OrderItemParameters `json:"order_items" binding:"required"`
}

func OrderCreate(c *gin.Context) {
	db := database.Connect(c)

	tokenData, found := tokenRepository.FindByAccessToken(db, helpers.GetAuthorizationToken(c))
	if found == 0 {
		c.JSON(400, gin.H{"status": 0, "message": "User is not found!"})
		return
	}

	var OrderParameter OrderParameters
	err := c.BindJSON(&OrderParameter)
	if err != nil {
		c.JSON(400, gin.H{"status": 0, "message": err.Error()})
		return
	}

	var OrderItemParameter []orderService.OrderItemParameter
	for _, item := range OrderParameter.OrderItems {
		OrderItemParameter = append(OrderItemParameter, orderService.OrderItemParameter{
			ProductsID: item.ProductsID,
			Qty:        item.Qty,
		})
	}

	newID, newErr := orderService.CreateOrder(c, orderService.OrderParameter{
		CustomersID:         tokenData.UsersID,
		ShippingVendor:      OrderParameter.ShippingVendor,
		ShippingPackage:     OrderParameter.ShippingPackage,
		RecipientName:       OrderParameter.RecipientName,
		RecipientAddress:    OrderParameter.RecipientAddress,
		RecipientCity:       OrderParameter.RecipientCity,
		RecipientPostalCode: OrderParameter.RecipientPostalCode,
		OrderItems:          OrderItemParameter,
	})

	// Send Email trigger
	type KafkaDt struct {
		Action string
		ID     uint
	}
	kafkaData, _ := json.Marshal(KafkaDt{"SEND_EMAIL_ORDER_SUCCESS", newID})
	kafkaErr := helpers.KafkaPushMessage("General", kafkaData)
	if kafkaErr != nil {
		c.JSON(500, gin.H{"status": 0, "message": kafkaErr.Error()})
		return
	}

	if newErr != nil {
		c.JSON(500, gin.H{"status": 0, "message": newErr.Error()})
	} else {
		c.JSON(200, gin.H{"status": 1, "message": "Thank you, your order has been created!"})
	}
}
