package handlers

import (
	"crypto/rand"
	Paymentgateways "ecom/PaymentGateways"
	"ecom/initializers"
	"ecom/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var Couponcheck models.Coupon

func Checkout(c *gin.Context) {
	var order models.Order
	var orderItems models.OrderItem
	var cart []models.Cart
	var address models.Address
	// var Payment models.Payment
	var Grandtotal int

	userid := c.GetUint("userID")
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(500, "failed to fetch data")
	}
	//========================== Cart Details ==========================================
	if err := initializers.DB.Joins("Product").Where("user_id = ?", userid).Find(&cart); err.Error != nil {
		c.JSON(500, "Failed to fetch data from Cart DB")
		fmt.Println("Failed to fetch data from Cart DB=====>", err.Error)
	}
	// ========================= Stock Check ============================================
	for _, view := range cart {

		quantity_price := int(view.Quantity) * view.Product.Price

		if int(view.Quantity) > view.Product.Quantity {
			c.JSON(500, "Product Out of Stock"+view.Product.Product_Name)
			return
		}
		Grandtotal += quantity_price
	}

	//========================= Cheking Coupon =========================================
	var coup string
	if err := initializers.DB.Where("code = ? AND coundition < ?", order.CouponCode,Grandtotal).First(&Couponcheck); err.Error != nil {
		fmt.Println("code===>",order.CouponCode,"condition====>",Couponcheck.Coundition,"Grandtotal=====>",Grandtotal,"couponcheck=====>",Couponcheck)
		c.JSON(500, "Invalid Coupon")
		coup = "No coupon added"
		fmt.Println("Invalid Coupon=====>", err.Error)
	} else {
		c.JSON(200, "Coupon Added")
			Grandtotal -= int(Couponcheck.Discount)
			coup = order.CouponCode
			fmt.Println("check==================>", Couponcheck.Code, order.CouponCode)
		}
		fmt.Println("coupon=====>", Couponcheck.Code, order.CouponCode)
	//========================== Address choosing ======================================
	if err := initializers.DB.Where("user_id = ? AND id = ?", userid, order.AddressID).First(&address); err.Error != nil {
		c.JSON(500, "Address not found")
		fmt.Println("Address not found=======>", err.Error)
		return
	}
	//========================= Creating Random OrderID ================================
	const charset = "123456789"
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println(err)
	}
	for i, b := range randomBytes {
		randomBytes[i] = charset[b%byte(len(charset))]
	}
	orderIdstring := string(randomBytes)
	orderId, _ := strconv.Atoi(orderIdstring)
	fmt.Println("------", orderId)
	//========================== Transaction ==============================================

	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	//======================== Payment Gateway ==============================================

	fmt.Println("orderid==>", orderId, "grandtotal==>", Grandtotal)
	if order.OrderPayment == "UPI" {
		OrderPaymentID, err := Paymentgateways.HandlePaymentSubmission(orderId, Grandtotal)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			tx.Rollback()
			return
		} else {
			c.JSON(200, gin.H{
				"message":   "Continue to Payment",
				"paymentID": OrderPaymentID,
			})
		}
		fmt.Println("orderpayment:==>", OrderPaymentID)
		fmt.Println("receipt====>", orderId)
		if err := initializers.DB.Create(&models.Payment{
			OrderID:       OrderPaymentID,
			Receipt:       orderId,
			PaymentStatus: "not done",
			PaymentAmount: Grandtotal,
		}); err.Error != nil {
			c.JSON(500, gin.H{"Error": "Failed to upload payment details"})
			fmt.Println("Failed to upload payment details", err.Error)
		}
	}
	//========================= Order Table management ====================================

	order = models.Order{
		ID:           uint(orderId),
		UserID:       int(userid),
		OrderPayment: order.OrderPayment,
		AddressID:    order.AddressID,
		CouponCode:   coup,
		OrderPrice:   Grandtotal,
		OrderDate:    time.Now(),
	}
	if err := tx.Create(&order); err.Error != nil {
		tx.Rollback()
		c.JSON(500, "Failed to Place Order")
		fmt.Println("Failed to Place Order=====>", err.Error)
		return
	}
	for _, view := range cart {
		subTotal := int(view.Quantity) * view.Product.Price
		orderItems = models.OrderItem{
			ProductID:     view.Product_Id,
			OrderID:       uint(orderId),
			OrderQuantity: int(view.Quantity),
			Subtotal:      float64(subTotal),
			Orderstatus:   "Pending",
		}
		fmt.Println("orderitems", orderItems)
		if err := tx.Create(&orderItems); err.Error != nil {
			tx.Rollback()
			c.JSON(500, "Failed to Place Order")
			fmt.Println("Failed to Place Order=====>", err.Error)
			return
		}
		//========================= Stock management =======================================
		view.Product.Quantity -= int(view.Quantity)
		if err := initializers.DB.Save(&view.Product); err.Error != nil {
			c.JSON(500, "Failed to update product stock")
			fmt.Println("Failed to update product stock======>", err.Error)
		} else {
			fmt.Println("Stock Updated=====>")
		}
	}

	// if err := initializers.DB.Where("user_id = ?", userid).Delete(&models.Cart{}); err.Error != nil {
	// 	c.JSON(500, "Failed to delete order")
	// 	return
	// }
	if err := tx.Commit(); err.Error != nil {
		tx.Rollback()
		c.JSON(500, "Failed to commit transaction")
		fmt.Println("Failed to commit transaction=====>", err.Error)
		return
	}
	if order.OrderPayment != "UPI" {
		c.JSON(200, gin.H{
			"message":      "Order Placed Succesfully",
			"Grand Total ": Grandtotal,
		})
	}

}
