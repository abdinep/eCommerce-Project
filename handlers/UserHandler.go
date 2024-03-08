package handlers

import (
	"ecom/initializers"
	"ecom/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var load []models.Product

func ProductLoadingPage(c *gin.Context) {

	if err := initializers.DB.Joins("Category").Find(&load); err.Error != nil {
		c.JSON(http.StatusBadRequest, "Failed to fetch product data")
		return
	}
	for _, view := range load {
		fmt.Println("products", view.ID, "===>", view.Product_Name, "===>", view.ImagePath1, "===>")
		c.JSON(http.StatusOK, gin.H{
			"ID":    view.ID,
			"Name":  view.Product_Name,
			"Image": view.ImagePath1,
		})
	}
}

func ProductDetails(c *gin.Context) {
	var product models.Product
	productID := c.Param("ID")
	if err := initializers.DB.First(&product, productID); err.Error != nil {
		c.JSON(http.StatusBadRequest, "failed to fetch product data")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"product Name":        product.Product_Name,
			"product image1":      product.ImagePath1,
			"product image2":      product.ImagePath2,
			"product image3":      product.ImagePath3,
			"product Price":       product.Price,
			"product size":        product.Size,
			"product discription": product.Description,
		})
	}
	if product.Quantity == 0 {
		c.JSON(http.StatusOK, gin.H{
			"stock status": "Out of stock",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stock status": "Item is available",
		})
	}

	rating := Ratingcalc(productID, c)

	c.JSON(http.StatusOK, gin.H{
		"rating": rating,
	})
	ReviewView(productID, c)

	for _, value := range load {
		if product.Category_id == value.Category_id && product.Product_Name != value.Product_Name {
			c.JSON(http.StatusOK, gin.H{
				"Showing":             "similar products",
				"product Name":        value.Product_Name,
				"product image":       value.ImagePath1,
				"product Price":       value.Price,
				"product size":        value.Size,
				"product discription": value.Description,
				"category":            value.Category.Name,
			})
		}
	}
}

//================================ END ================================================
//=============================== Rating ==============================================

func RatingStrore(c *gin.Context) {
	var userrate models.Rating
	var dbrate models.Rating
	// ID := c.Param("ID")
	if err := c.ShouldBindJSON(&userrate); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to bind data")
	}
	result := initializers.DB.First(&dbrate, "product_id=?", userrate.ProductId)
	if result.Error != nil {
		userrate.Value = 1
		if err := initializers.DB.Create(&userrate).Error; err != nil {
			c.JSON(http.StatusBadRequest, "failed to store")
		} else {
			c.JSON(http.StatusOK, "Thanks for rating")
		}
	} else {
		err := initializers.DB.Model(&dbrate).Where("product_id=?", userrate.ProductId).Updates(models.Rating{
			Users: dbrate.Users + 1,
			Value: dbrate.Value + userrate.Value,
		})
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, "failed to update data")
		} else {
			c.JSON(http.StatusOK, "Thanks for rating")
		}
	}
	dbrate = models.Rating{}

}
func Ratingcalc(id string, c *gin.Context) float64 {
	var ratinguser models.Rating
	if err := initializers.DB.First(&ratinguser, "product_id=?", id); err.Error != nil {
		c.JSON(http.StatusBadRequest, "failed to fetch data")
	} else {
		averageratio := float64(ratinguser.Value) / float64(ratinguser.Users)
		ratinguser = models.Rating{}
		result := fmt.Sprintf("%.1f", averageratio)
		averageratio, _ = strconv.ParseFloat(result, 64)
		return averageratio
	}
	return 0
}

// ========================================= END ==================================================
// ================================== Review ======================================================
func ReviewStore(c *gin.Context) {
	var reviewstore models.Review
	if err := c.ShouldBindJSON(&reviewstore); err != nil {
		c.JSON(http.StatusBadRequest, "failed to bind data")
	} else {
		reviewstore.Time = time.Now().Format("2006-01-02")
		if err := initializers.DB.Create(&reviewstore); err.Error != nil {
			c.JSON(http.StatusBadRequest, "failed to store review")
		} else {
			c.JSON(http.StatusOK, "Thank you for your valuable feedback")
		}
	}
}
func ReviewView(id string, c *gin.Context) {
	var reviewView []models.Review
	if err := initializers.DB.Joins("User").Find(&reviewView).Where("product_id=?", id); err.Error != nil {
		c.JSON(http.StatusBadRequest, "failed to fetch reviews")

	} else {
		productId, _ := strconv.Atoi(id)
		for _, val := range reviewView {
			if val.ProductId == uint(productId) {
				c.JSON(http.StatusOK, gin.H{
					"review_user": val.User.Name,
					"review_date": val.Time,
					"review":      val.Review,
				})
			}
		}
	}

}

// ========================================= END ==================================================
// ================================ Address management ============================================
func Add_Address(c *gin.Context) {
	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(500, "failed to fetch data")
	}
	if err := initializers.DB.Create(&address); err.Error != nil {
		c.JSON(500, "Existing Address")
	} else {
		c.JSON(http.StatusOK, "New Address added")
	}
}
func Edit_Address(c *gin.Context) {
	var address models.Address
	id := c.Param("ID")
	if err := initializers.DB.First(&address, id); err.Error != nil {
		c.JSON(500, "Failed to fetch address")
	} else {
		if err := c.ShouldBindJSON(&address); err != nil {
			c.JSON(500, "failed to bind address")
			return
		}
		if err := initializers.DB.Save(&address); err.Error != nil {
			c.JSON(500, "failed to edit address")
		} else {
			c.JSON(200, "Updated Address")
		}
	}

}
func Delete_Address(c *gin.Context) {
	var address models.Address
	id := c.Param("ID")
	if err := initializers.DB.First(&address, id); err.Error != nil {
		c.JSON(500, "failed to fetch data")
	} else {
		if err := initializers.DB.Delete(&address); err.Error != nil {
			c.JSON(500, "Address cant be deleted")
		} else {
			c.JSON(200, "Address Deleted successfully")
		}
	}

}

// ========================================= END ==================================================
// =================================== User Profile ===============================================
func User_Details(c *gin.Context) {
	var details models.User
	id := c.Param("ID")
	if err := initializers.DB.First(&details, id); err.Error != nil {
		c.JSON(500, "Failed to fetch data")
		fmt.Println("Error", err.Error)
	} else {
		c.JSON(200, gin.H{
			"user_name":   details.Name,
			"user_email":  details.Email,
			"user_mobile": details.Mobile,
			"user_gender": details.Gender,
			"user_status": details.Status,
			// "user_address" : details.

		})
	}
}
func Edit_Profile(c *gin.Context) {
	var edit models.User
	id := c.Param("ID")
	if err := initializers.DB.First(&edit, id); err.Error != nil {
		c.JSON(500, "Failed to fetch data from DB")
		fmt.Println("Failed to fetch data from DB=====>", err.Error)
	} else {
		if err := c.ShouldBindJSON(&edit); err != nil {
			c.JSON(500, "failed to bind profile details")
			return
		}
		if err := initializers.DB.Save(&edit); err.Error != nil {
			c.JSON(500, "failed to edit details")
			fmt.Println("failed to edit details", err.Error)
		} else {
			c.JSON(200, "Updated Profile details")
		}
	}

}
func View_Address(c *gin.Context) {
	var address []models.Address
	id := c.Param("ID")
	if err := initializers.DB.Find(&address).Where("UserId = ?", id); err.Error != nil {
		c.JSON(500, "Failed to find address")
		fmt.Println(err.Error, address)
	} else {
		userID, _ := strconv.Atoi(id)
		for _, view := range address {
			if view.UserId == userID {

				c.JSON(http.StatusOK, gin.H{
					"Address_Type": view.Type,
					"Address_ID":   view.ID,
					"User_Address": view.Address,
					"User_City":    view.City,
					"User_State":   view.State,
					"User_Pincode": view.Pincode,
					"User_Country": view.Country,
					"User_Phone":   view.Phone,
				})
			}
		}
	}
}
func View_Orders(c *gin.Context) {
	var order []models.Order

	userID := c.Param("ID")
	if err := initializers.DB.Joins("Product").Where("user_id = ?", userID).Find(&order); err.Error != nil {
		c.JSON(500, "Currently no Orders")
		fmt.Println("Currently no Orders========>", err.Error)

	} else {
		count := 0
		for _, view := range order {
			c.JSON(200, gin.H{
				"Order_ID":         view.ID,
				"Product_Name":     view.Product.Product_Name,
				"Selected_Address": view.AddressID,
				"Applied_Coupon":   view.Coupon_Code,
				"Order_Quantity":   view.Order_Quantity,
				"Order_Price":      view.Order_Price,
				"Payment_Method":   view.Order_Payment,
				"order_status":     view.Order_status,
			})
			count += 1
		}
		c.JSON(200, gin.H{
			"No.Order": count,
		})
	}
}
func Cancel_Orders(c *gin.Context) {
	var order models.Order
	var cancel models.Order
	var quantity models.Product
	userid := c.Param("ID")
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(500, "Failed to bind ")
		return
	}

	initializers.DB.Where("product_id = ?", order.ProductID).First(&quantity)

	if err := initializers.DB.Where("user_id = ? AND product_id = ?", userid, order.ProductID).First(&cancel); err.Error != nil {
		c.JSON(500, "Order not exist")
		fmt.Println("Order not exist", err.Error)
	} else {
		cancel.Order_status = "Order Canceled"
		if err := initializers.DB.Save(&cancel); err.Error != nil {
			c.JSON(500, "Failed to cancel your order")
			fmt.Println("Failed to cancel your order", err.Error)
		} else {
			c.JSON(200, "Order canceled successfully")
			quantity.Quantity += cancel.Order_Quantity
			initializers.DB.Save(&quantity)
			fmt.Println("++++++++++", quantity.Quantity, cancel.Order_Quantity, "+++++++++++")
		}
	}
}

// ========================================= END ==================================================
