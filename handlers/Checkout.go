package handlers

import (
	"ecom/initializers"
	"ecom/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Checkout(c *gin.Context) {
	var order models.Order
	var couponcheck models.Coupon
	var cart []models.Cart
	var address models.Address
	var Grandtotal int

	userid := c.Param("ID")
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(500, "failed to fetch data")
	}
	//========================== Cart Details ==========================================
	if err := initializers.DB.Joins("Product").Where("user_id = ?", userid).Find(&cart); err.Error != nil {
		c.JSON(500, "Failed to fetch data from Cart DB")
		fmt.Println("Failed to fetch data from Cart DB=====>", err.Error)
	}
	//========================= Cheking Coupon =========================================
	if err := initializers.DB.Where("code = ? ", order.Coupon_Code).First(&couponcheck); err.Error != nil {
		c.JSON(500, "Invalid Coupon")
		fmt.Println("Invalid Coupon=====>", err.Error)
		return
	} else {
		c.JSON(200, "Coupon discount Added")
		fmt.Println("coupon=====>", couponcheck.Code, order.Coupon_Code)

	}
	//========================== Address choosing ======================================
	if err := initializers.DB.Where("user_id = ? AND id = ?", userid, order.AddressID).First(&address); err.Error != nil {
		c.JSON(500, "Address not found")
		fmt.Println("Address not found=======>", err.Error)
		return
	}

	id, _ := strconv.Atoi(userid)
	count := 1
	for _, view := range cart {

		quantity_price := int(view.Quantity) * view.Product.Price

		if int(view.Quantity) > view.Product.Quantity {
			c.JSON(500, "Product Out of Stock"+view.Product.Product_Name)
			return
		}
		view.Product.Quantity -= int(view.Quantity)
		if err := initializers.DB.Save(&view.Product); err.Error != nil {
			c.JSON(500, "Failed to update product stock")
			fmt.Println("Failed to update product stock======>", err.Error)
		} else {
			fmt.Println("Stock Updated=====>")
		}
		var coup string
		if couponcheck.Code == order.Coupon_Code && count != 0 {
			quantity_price = quantity_price - int(couponcheck.Discount)
			coup = order.Coupon_Code
			count--
			fmt.Println("check==================>", couponcheck.Code, order.Coupon_Code)
		} else {
			coup = "No coupon added"
		}

		order.Order_Price = quantity_price
		order = models.Order{
			UserID:         id,
			Order_Payment:  order.Order_Payment,
			AddressID:      order.AddressID,
			ProductID:      view.Product_Id,
			Order_Quantity: int(view.Quantity),
			Coupon_Code:    coup,
			Order_Price:    order.Order_Price,
			Order_Date:     time.Now(),
			Order_status: "Pending",
		}

		if err := initializers.DB.Create(&order); err.Error != nil {
			c.JSON(500, "Failed to Place Order")
			fmt.Println("Failed to Place Order", err.Error)
		}
		Grandtotal += quantity_price
	}
	c.JSON(200, gin.H{
		"message":      "Order Placed Succesfully",
		"Grand Total ": Grandtotal,
	})

	if err := initializers.DB.Where("user_id = ?", userid).Delete(&models.Cart{}); err.Error != nil {
		c.JSON(500, "Failed to delete order")
		return
	}
}
