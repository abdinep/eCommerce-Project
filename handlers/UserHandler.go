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
	if product.Quantity == 0 {
		c.JSON(http.StatusOK,gin.H{
			"stock status": "Out of stock",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"stock status": "Item is available",
		})
	}
	
	rating := Ratingcalc(productID,c)

	c.JSON(http.StatusOK,gin.H{
		"rating":rating,
	})
	ReviewView(productID,c)
	
	for _,value:=range load{
		if product.Category_id == value.Category_id && product.Product_Name != value.Product_Name{
			c.JSON(http.StatusOK,gin.H{
				"Showing":	"similar products",
				"product Name": value.Product_Name,
				"product image":value.ImagePath1,
				"product Price":value.Price,
				"product size": value.Size,
				"product discription":value.Description,
			})
		}
	}
}
//================================ END ================================================
//=============================== Rating ==============================================

func RatingStrore(c *gin.Context){
	var userrate models.Rating
	var dbrate models.Rating
	// ID := c.Param("ID")
	if err:=c.ShouldBindJSON(&userrate);err != nil{
		c.JSON(http.StatusBadRequest,"Failed to bind data")
	}
	result := initializers.DB.First(&dbrate,"product_id=?",userrate.ProductId)
	if result.Error != nil{
		userrate.Value = 1
		if err := initializers.DB.Create(&userrate).Error; err != nil{
			c.JSON(http.StatusBadRequest,"failed to store")
		}else{
			c.JSON(http.StatusOK,"Thanks for rating")
		}
	}else{
		err := initializers.DB.Model(&dbrate).Where("product_id=?",userrate.ProductId).Updates(models.Rating{
			Users: dbrate.Users + 1,
			Value: dbrate.Value + userrate.Value,
		})
		if err.Error != nil{
			c.JSON(http.StatusBadRequest,"failed to update data")
		}else{
			c.JSON(http.StatusOK,"Thanks for rating")
		}
	}
	dbrate = models.Rating{}

}
func Ratingcalc(id string,c *gin.Context) float64{
	var ratinguser models.Rating
	if err := initializers.DB.First(&ratinguser,"product_id=?",id);err.Error != nil{
		c.JSON(http.StatusBadRequest,"failed to fetch data")
	}else{
		averageratio := float64(ratinguser.Value) / float64(ratinguser.Users)
		ratinguser = models.Rating{}
		result := fmt.Sprintf("%.1f",averageratio)
		averageratio,_ = strconv.ParseFloat(result,64)
		return averageratio
	}
	return 0
}
//========================================= END ==================================================
//================================== Review ======================================================
func ReviewStore(c *gin.Context){
	var reviewstore models.Review
	if err := c.ShouldBindJSON(&reviewstore); err != nil{
		c.JSON(http.StatusBadRequest,"failed to bind data")
	}else{
		reviewstore.Time = time.Now().Format("2006-01-02")
		if err := initializers.DB.Create(&reviewstore); err.Error != nil{
			c.JSON(http.StatusBadRequest,"failed to store review")
		}else{
			c.JSON(http.StatusOK,"Thank you for your valuable feedback")
		}
	}
}
func ReviewView(id string, c *gin.Context){
	var reviewView []models.Review
	if err := initializers.DB.Joins("User").Find(&reviewView).Where("product_id=?",id); err.Error != nil{
		c.JSON(http.StatusBadRequest,"failed to fetch reviews")
		
	}else{
		productId,_:= strconv.Atoi(id)
		for _, val := range reviewView{
			if val.ProductId == uint(productId){
				c.JSON(http.StatusOK,gin.H{
					"review_user" : val.User.Name,
					"review_date" : val.Time,
					"review": val.Review,
				})
			}else{
				c.JSON(http.StatusOK,gin.H{
					"review" : "No review",
				})
			}
		}
	}
	
}
//========================================= END ==================================================