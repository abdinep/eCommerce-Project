package handlers

import (
	"ecom/initializers"
	"ecom/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var product models.Product

// ====================== Display all user details in admin panel ====================================
func List_user(c *gin.Context) {
	var list []models.User
	initializers.DB.Find(&list)
	for _, data := range list {
		c.JSON(200, gin.H{
			"ID":     data.ID,
			"User":   data.Name,
			"Email":  data.Email,
			"Mobile": data.Mobile,
			"status": data.Status,
		})
	}
}

//=========================== END ==================================

// ======================= Adding products to the DB ===============================
func Add_Product(c *gin.Context) {

	var cat_id models.Category
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(501, "failed to bind json")
	}
	fmt.Println("============", product, "==================")
	result := initializers.DB.First(&cat_id, product.Category_id)
	fmt.Println("============", cat_id, "=====================")
	if result.Error != nil {
		c.JSON(404, "Category not found")
	} else {

		c.JSON(200, "Upload Product Images ")
	}
}

// ==================================== END =========================================
// ================================= Upload Product Image ===========================
func ProductImage(c *gin.Context) {
	file, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to fetch images")
	}
	files := file.File["images"]
	var imagepaths []string

	for _, val := range files {
		filepath := "./images/" + val.Filename
		if err = c.SaveUploadedFile(val, filepath); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to save images")
		}
		imagepaths = append(imagepaths, filepath)
	}
	product.ImagePath1 = imagepaths[0]
	product.ImagePath2 = imagepaths[1]
	product.ImagePath3 = imagepaths[2]
	upload := initializers.DB.Create(&product)
	if upload.Error != nil {
		c.JSON(501, "Product already exist")
		return
	} else {
		c.JSON(200, "Product added successfully")
	}
	product = models.Product{}

}

// ==================================== END =========================================

//======================= Category Adding to the DB ================================

func Category(c *gin.Context) {
	var cat models.Category
	c.ShouldBindJSON(&cat)
	upload := initializers.DB.Create(&cat)
	if upload.Error != nil {
		c.JSON(404, "failed to upload category")
	} else {
		c.JSON(200, "New Category added")
	}
}

//==================================== END ===========================================

func View_Category(c *gin.Context) {
	var View []models.Category
	//  var checkcategory models.Categories
	initializers.DB.Find(&View)

	// initializers.DB.First(&checkcategory,"",View.Category_id)
	for _, view := range View {
		c.JSON(200, gin.H{
			"ID":            view.ID,
			"category Name": view.Name,
			"Discription":   view.Description,
			"status":        view.Status,
		})
	}
}

// ==================================== Editing Category ===========================================
func Edit_Category(c *gin.Context) {
	var edit models.Category
	id := c.Param("ID")
	result := initializers.DB.First(&edit, id)
	fmt.Println("(===============", edit, "===========)(", id, "===================)")
	if result.Error != nil {
		c.JSON(501, "Category not found")
	} else {
		err := c.ShouldBindJSON(&edit)
		if err != nil {
			c.JSON(501, "failed to bind json")
		}
		save := initializers.DB.Save(&edit)
		if save.Error != nil {
			c.JSON(501, "Failed to update Category")
		} else {
			c.JSON(200, "Category updated successfully")
		}
	}
}

//==================================== END ===========================================
//==================================== Deleting Categories ===========================================

func Delete_Category(c *gin.Context) {
	var delete models.Category
	cat := c.Param("ID")
	err := initializers.DB.First(&delete, cat)
	if err.Error != nil {
		c.JSON(501, "Category cant be deleted")
	} else {
		initializers.DB.Delete(&delete)
		c.JSON(200, "Category Deleted")
	}
}

//==================================== END ===========================================

// ====================== Showing all products in admin page ==========================
func View_Product(c *gin.Context) {
	var View []models.Product
	//  var checkcategory models.Categories
	initializers.DB.Where("deleted_at IS NULL").Preload("Category").Find(&View)
	for _, view := range View {
		c.JSON(200, gin.H{
			"product ID":       view.ID,
			"product Name":     view.Product_Name,
			"product quantity": view.Quantity,
			"product price":    view.Price,
			"product size":     view.Size,
			"Category":         view.Category.Name,
			"product image1":   view.ImagePath1,
			"Stock":            view.Quantity,
		})
	}
}

//==================================== END ========================================

// ====================== Editing user detailes in admin panel =========================
func Edit_User(c *gin.Context) {
	var edit models.User
	user := c.Param("ID")
	result := initializers.DB.First(&edit, user)
	fmt.Println("(===============", edit, "===========)(", user, "===================)")
	if result.Error != nil {
		c.JSON(500, "User not found")
	} else {

		err := c.ShouldBindJSON(&edit)
		if err != nil {
			c.JSON(500, "Failed to bind json")
		}
		error := initializers.DB.Save(&edit)
		if error.Error != nil {
			c.JSON(500, "failed to update user")
		} else {
			c.JSON(200, "User updated successfully")
		}
	}

}

//========================== END =====================================

// =========================== Editing product detailes in admin panel =========================
func Edit_Product(c *gin.Context) {
	var edit models.Product
	product := c.Param("ID")
	result := initializers.DB.First(&edit, product)
	fmt.Println("(===============", edit, "===========)(", product, "===================)")
	if result.Error != nil {
		c.JSON(501, "Product not found")
	} else {
		err := c.ShouldBindJSON(&edit)
		if err != nil {
			c.JSON(501, "failed to bind json")
		}
		save := initializers.DB.Save(&edit)
		if save.Error != nil {
			c.JSON(501, "Failed to update Product")
		} else {
			c.JSON(200, "Product updated successfully")
		}
	}
}

//========================= END ================================

// =================== Deleting products in admin panel =================================
func Delete_Product(c *gin.Context) {
	var delete models.Product
	product := c.Param("ID")
	err := initializers.DB.First(&delete, product)
	if err.Error != nil {
		c.JSON(501, "Product cant be deleted")
	} else {
		initializers.DB.Delete(&delete)
		c.JSON(200, "Product Deleted")
	}
}

//========================== END ===========================================

// =========================== User Block/Unblock in admin panel ===========================
func Status(c *gin.Context) {
	var check models.User
	user := c.Param("ID")
	initializers.DB.First(&check, user)
	if check.Status == "Active" {
		initializers.DB.Model(&check).Update("status", "Blocked")
		c.JSON(200, "User Blocked")
	} else {
		initializers.DB.Model(&check).Update("status", "Active")
		c.JSON(200, "User Unblocked")
	}

}

//=================================== END =====================================
//============================= Coupon Management =============================

func Coupon(c *gin.Context) {
	var coupon models.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(500, "Failed to fetch data")
	} else {
		if err := initializers.DB.Create(&coupon); err.Error != nil {
			c.JSON(500, "Coupon already exist")
			fmt.Println("Coupon already exist", err.Error)
		} else {
			c.JSON(200, "New Coupon added")
		}
	}
}

// ================================== END =======================================
// ========================== Order management ==================================
func Admin_View_order(c *gin.Context) {
	var order []models.Order
	count := 0
	if err := initializers.DB.Find(&order); err.Error != nil {
		c.JSON(500, "Failed to fetch order")
		return
	}
	for _, view := range order {
		c.JSON(200, gin.H{
			"OrderID":         view.ID,
			"userid":          view.UserID,
			"SelectedAddress": view.AddressID,
			"AppliedCoupon":   view.CouponCode,
			"OrderPrice":      view.OrderPrice,
			"PaymentMethod":   view.OrderPayment,
		})
		count += 1
		c.JSON(200,gin.H{"No.Orders":count})
	}
}
func ViewOrderDetails(c *gin.Context){
	var orderitem []models.OrderItem
	orderid := c.Param("ID")
	if err := initializers.DB.Where("order_id = ?",orderid).Joins("Product").Find(&orderitem); err.Error != nil{
		c.JSON(500,gin.H{"Error":"Produt not found"})
	}else{
		for _,view := range orderitem{
			subTotal := view.OrderQuantity * view.Product.Price
			c.JSON(200,gin.H{
				"OrdeD":view.OrderID,
				"OrderItemsID":    view.ID,
				"ProductName":     view.Product.Product_Name,
				"SelectedAddress": view.Order.AddressID,
				"AppliedCoupon":   view.Order.CouponCode,
				"OrderQuantity":   view.OrderQuantity,
				"OrderPrice":      subTotal,
				"PaymentMethod":   view.Order.OrderPayment,
				"orderStatus":     view.Orderstatus,
			})
		}
	}
}
func ChangeOrderStatus(c *gin.Context) {
	var order models.OrderItem
	var update models.OrderItem
	orderid := c.Param("ID")
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(500, gin.H{"Error":"Add status"})
		return
	}
	if err := initializers.DB.First(&update, orderid); err.Error != nil {
		c.JSON(500, gin.H{"Error":"Order not found"})
		fmt.Println("Order not found======>", err.Error)
		return
	}
	fmt.Println("========>", order.Orderstatus)
	update.Orderstatus = order.Orderstatus
	initializers.DB.Save(&update)
	c.JSON(200, gin.H{"Message":"Order status changed"})
}
