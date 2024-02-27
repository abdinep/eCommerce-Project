package handlers

import (
	"ecom/initializers"
	"ecom/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
var load []models.Product
func ProductLoadingPage(c *gin.Context){
	
	if err:=initializers.DB.Joins("Category").Find(&load);err.Error != nil{
		c.JSON(http.StatusBadRequest,"Failed to fetch product data")
	}
	for _,view := range load{
		fmt.Println("products",view.ID,"===>",view.Product_Name,"===>",view.ImagePath1,"===>")
		c.JSON(http.StatusOK,gin.H{
			"ID"	: view.ID,
			"Name"	: view.Product_Name,
			"Image"	: view.ImagePath1,
		})
	}
}
func ProductDetails(c *gin.Context){
	var product models.Product
	productID := c.Param("ID")
	if err:=initializers.DB.First(&product,productID);err.Error != nil{
		c.JSON(http.StatusBadRequest,"failed to fetch product data")
	}else{
		c.JSON(http.StatusOK,gin.H{
			"product Name": product.Product_Name,
			"product image1":product.ImagePath1,
			"product image2":product.ImagePath2,
			"product image3":product.ImagePath3,
			"product Price":product.Price,
			"product size": product.Size,
			"product discription":product.Description,
		})
	}
	for _,value:=range load{
		if product.Category_id == value.Category_id && product.Product_Name != value.Product_Name{
			c.JSON(http.StatusOK,gin.H{
				"Showing":	"similar products",
				"product Name": value.Product_Name,
				"product image1":value.ImagePath1,
				"product image2":value.ImagePath2,
				"product image3":value.ImagePath3,
				"product Price":value.Price,
				"product size": value.Size,
				"product discription":value.Description,
			})
		}
	}
}
