package handlers

import (
	"ecom/initializers"
	"ecom/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add_Cart(c *gin.Context) {
	var cart models.Cart
	var dbcart models.Cart

	id := c.Param("ID")
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(500, "failed to bind data")
		log.Fatal("failed to bind data", err)
	} else {
		Product_Id, _ := strconv.Atoi(id)
		cart.Product_Id = Product_Id
		initializers.DB.First(&dbcart,"user_id=?",cart.User_id )
		fmt.Println("dbcart=====>", dbcart)
		if cart.Product_Id == dbcart.Product_Id {
			cart.Quantity += 1
			if err := initializers.DB.Model(&dbcart).Where("user_id = ? AND product_id = ?",dbcart.User_id,dbcart.Product_Id).Updates(models.Cart{Quantity :cart.Quantity}); err.Error != nil {
				c.JSON(500, "failed to update data")
				log.Fatal("failed to update data", err.Error)
			} else {
				c.JSON(200, "Added one more quantity")
			}
		}else{
			if err := initializers.DB.Create(&cart); err.Error != nil {
				c.JSON(500, "failed to add to cart")
				log.Fatal("failed to add to cart", err)
			} else {
				c.JSON(200, "New product added to your cart")
			}
		}
	}
	
}
func View_Cart(c *gin.Context) {
	var cart []models.Cart
	// var userId models.Cart
	id := c.Param("ID")
	if err := initializers.DB.Joins("Product").Find(&cart).Where("user_id = ?",id); err.Error != nil {
		c.JSON(500, "Product not found")
		log.Fatal("product not found", err.Error)
	} else {
		count := 0
		for _, view := range cart {
			c.JSON(200, gin.H{
				"product_id": view.Product_Id,
				"Product_Name":     view.Product.Product_Name,
				"Product_image":    view.Product.ImagePath1,
				"Product_Price":    view.Product.Price,
				"Product_Quantity": view.Quantity,
				
			})
			 count += 1
		}
		c.JSON(200,gin.H{
			"total_number" : count, 
		})
	}
}
