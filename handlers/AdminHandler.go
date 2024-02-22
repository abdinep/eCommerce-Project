package handlers

import (
	"ecom/initializers"
	"ecom/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
		})
	}
}

//=========================== END ==================================

// ======================= Adding products to the DB ===============================
func Add_Product(c *gin.Context) {
	var product models.Product
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

		// product.Category_id = int(cat_id.ID)
		fmt.Println("==============", product.Category_id, "==================")
		upload := initializers.DB.Create(&product)
		if upload.Error != nil {
			c.JSON(501, "Product already exist")
			return
		}
		c.JSON(200, "New Product Added")
	}
}

//==================================== END =========================================

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
		c.JSON(501, "User not found")
	} else {

		err := c.ShouldBindJSON(&edit)
		if err != nil {
			c.JSON(501, "Failed to bind json")
		}
		error := initializers.DB.Save(&edit)
		if error.Error != nil {
			c.JSON(501, "failed to update user")
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
