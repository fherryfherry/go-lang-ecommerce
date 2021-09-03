package orderService

import (
	"ecommerce/database"
	"ecommerce/helpers"
	"ecommerce/models"
	"ecommerce/repositories/productRepository"
	"ecommerce/repositories/shippingVendorRepository"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type OrderItemParameter struct {
	ProductsID uint
	Qty        float32
}

type OrderParameter struct {
	OrderNo             string
	CustomersID         int
	ShippingVendor      string
	ShippingPackage     string
	ShippingPrice       float32
	RecipientName       string
	RecipientAddress    string
	RecipientPostalCode string
	RecipientCity       string
	TotalWeight         float32
	GrandTotal          float32
	PaymentStatus       string
	DeliveryStatus      string
	OrderStatus         string
	OrderItems          []OrderItemParameter
}

func MakeOrderNo(db *gorm.DB) string {
	var OrderNo string
	var Order models.Orders
	db.Last(&Order)
	OrderNo = helpers.PadLeft(int64(Order.ID+1), 6)
	OrderNo = "INV" + OrderNo
	return OrderNo
}

func CreateOrder(c *gin.Context, Param OrderParameter) error {
	db := database.Connect(c)

	err := db.Transaction(func(tx *gorm.DB) error {
		// Check shipping price
		shippingPrice, _ := shippingVendorRepository.FindPrice(db, "Jakarta", Param.RecipientCity, Param.TotalWeight)

		// Create a header order
		Order := new(models.Orders)
		Order.OrderNo = MakeOrderNo(tx)
		Order.CustomersID = Param.CustomersID
		Order.ShippingVendor = Param.ShippingVendor
		Order.ShippingPackage = Param.ShippingPackage
		Order.ShippingPrice = shippingPrice
		Order.RecipientName = Param.RecipientName
		Order.RecipientAddress = Param.RecipientAddress
		Order.RecipientCity = Param.RecipientCity
		Order.RecipientPostalCode = Param.RecipientPostalCode
		Order.PaymentStatus = "Waiting Payment"
		Order.DeliveryStatus = "Pending"
		Order.OrderStatus = "New Order"

		if err := tx.Create(&Order).Error; err != nil {
			return err
		}

		var grandTotal float32

		// Save the order items
		for _, item := range Param.OrderItems {
			Product, found := productRepository.FindById(db, item.ProductsID)
			if found == 0 {
				return errors.New("Product ID " + strconv.Itoa(int(item.ProductsID)) + " not found")
			}

			var subTotal float32
			subTotal = Product.Price * item.Qty
			OrderItem := new(models.OrderItems)
			OrderItem.OrdersID = Order.ID
			OrderItem.ProductsID = item.ProductsID
			OrderItem.Price = Product.Price
			OrderItem.Qty = item.Qty
			OrderItem.SubTotal = subTotal
			if err := tx.Create(&OrderItem).Error; err != nil {
				return err
			}

			grandTotal += subTotal
		}

		Order.GrandTotal = grandTotal
		if err := tx.Save(&Order).Error; err != nil {
			return err
		}

		// Commit all transactions
		return nil
	})

	return err
}
