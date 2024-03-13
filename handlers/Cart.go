package handlers

import (
	"ecom/initializers"
	"ecom/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add_Cart(c *gin.Context) {
	var cart models.Cart
	var dbcart models.Cart

	id := c.Param("ID")
	Product_Id, _ := strconv.Atoi(id)
	cart.Product_Id = Product_Id
	cart.User_id = int(c.GetUint("userID"))
	initializers.DB.First(&dbcart, "product_id=? AND user_id = ?", cart.Product_Id, cart.User_id)
	fmt.Println("dbcart=====>", dbcart)
	if cart.Product_Id == dbcart.Product_Id {
		c.JSON(500, "Product already added to Cart")
	} else {
		cart.Quantity = 1
		if err := initializers.DB.Create(&cart); err.Error != nil {
			c.JSON(500, "failed to add to cart")
			fmt.Println("failed to add to cart=====>", err)
		} else {
			c.JSON(200, "New product added to your cart")

		}
	}
}
func Add_Quantity_Cart(c *gin.Context) {
	var add models.Cart
	// var userid models.Cart
	var Product models.Product
	id := c.Param("ID")
	userid := c.GetUint("userID")
	if err := initializers.DB.First(&Product, id); err.Error != nil {
		c.JSON(500, "Failed to fetch data from product DB")
		fmt.Println("Failed to fetch data from product DB", err.Error)
		return
	}
	if err := initializers.DB.Where("product_id = ? AND user_id = ? ", id, userid).First(&add); err.Error != nil {
		c.JSON(500, "Failed to fetch data")
		fmt.Println("Failed to fetch data=====>", err.Error)
	} else {
		add.Quantity += 1
		if add.Quantity <= 5 && add.Quantity <= uint(Product.Quantity) {
			c.JSON(200, gin.H{"count": add.Quantity})
			if err := initializers.DB.Model(&add).Updates(models.Cart{Quantity: add.Quantity}); err.Error != nil {
				c.JSON(500, "Failed to update quantity ")
				fmt.Println("Failed to update quantity", err.Error)
			} else {
				c.JSON(200, "Added one more quantity")
			}
		} else {
			c.JSON(500, "Cant add more quantity or No stock")
		}
	}
}
func Remove_Quantity_cart(c *gin.Context) {
	// var remove models.Cart
	var dbremove models.Cart
	id := c.Param("ID")
	userid := c.GetUint("userID")
	if err := initializers.DB.Where("product_id = ? AND user_id = ? ", id, userid).First(&dbremove); err.Error != nil {
		c.JSON(500, "Failed to fetch data")
		fmt.Println("Failed to fetch data=====>", err.Error,userid,id)
	} else {
		dbremove.Quantity -= 1
		c.JSON(200, gin.H{"count": dbremove.Quantity})
		if err := initializers.DB.Model(&dbremove).Updates(models.Cart{Quantity: dbremove.Quantity}); err.Error != nil {
			c.JSON(500, "Failed to update quantity")
			fmt.Println("Failed to update quantity", err.Error)
		} else {
			c.JSON(200, "removed one more quantity")
		}
	}
}
func View_Cart(c *gin.Context) {
	var cart []models.Cart
	var quantity_price int
	var Grandtotal = 0
	id := c.GetUint("userID")
	if err := initializers.DB.Joins("Product").Find(&cart).Where("user_id = ?", id); err.Error != nil {
		c.JSON(500, "Product not found")
		fmt.Println("product not found=====>", err.Error)
	} else {

		count := 0
		// id, _ := strconv.Atoi(id)
		for _, view := range cart {
			if view.User_id == int(id) {
				quantity_price = int(view.Quantity) * view.Product.Price
				c.JSON(200, gin.H{
					"product_id":       view.Product_Id,
					"Product_Name":     view.Product.Product_Name,
					"Product_image":    view.Product.ImagePath1,
					"Product_Price":    view.Product.Price,
					"Product_Quantity": view.Quantity,
				})
				Grandtotal += quantity_price
				count += 1
			}
		}
		c.JSON(200, gin.H{
			"total_number": count,
			"Grand Total":  Grandtotal,
		})
	}
}
func Remove_Cart_Product(c *gin.Context) {
	var remove models.Cart
	id := c.Param("ID")
	userid := c.GetUint("userID")
	if err := initializers.DB.Where("product_id = ? AND user_id = ? ", id, userid).First(&remove); err.Error != nil {
		c.JSON(500, "Failed to fetch data")
		fmt.Println("Failed to fetch data=====>", err.Error)
	} else {
		if err := initializers.DB.Delete(&remove); err.Error != nil {
			c.JSON(500, "Cant delete the product")
			fmt.Println("Cant delete the product=====>", err.Error)
		} else {
			c.JSON(200, "Prouct removed from cart")
		}
	}
}
