package handlers

import (
	"ecom/initializers"
	"ecom/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddOffer(c *gin.Context) {
	var offer models.Offer
	var product models.Product

	productid := c.Param("ID")
	c.ShouldBindJSON(&offer)
	if err := initializers.DB.Where("id = ?", productid).First(&product); err.Error != nil {
		c.JSON(500, gin.H{"Error": "Product not available"})
		fmt.Println("Product not available======>", err.Error)
	} else {
		offer.Created = time.Now()
		offer.ProductId, _ = strconv.Atoi(productid)
		if err := initializers.DB.Create(&offer); err.Error != nil {
			c.JSON(500, gin.H{"Error": "Failed to add offer"})
			fmt.Println("Failed to add offer=====>", err.Error)
		} else {
			c.JSON(200, gin.H{"Message": "Offer Added for the Product"})
		}
	}
}
func ViewOffer(c *gin.Context) {
	var offer []models.Offer

	if err := initializers.DB.Joins("Product").Find(&offer); err.Error != nil {
		c.JSON(500, gin.H{"Error": "Offer not found"})
	} else {
		for _, view := range offer {
			c.JSON(200, gin.H{
				"ProductID":   view.ProductId,
				"ProductName": view.Product.Product_Name,
				"OfferName":   view.OfferName,
				"OfferAmount": view.Amount,
				"Created":     view.Created,
				"Expire":      view.Expire,
			})
		}
	}

}

func OfferCalc(productid int, c *gin.Context) float64 {
	var offercheck models.Offer
	var Discount float64
	if err := initializers.DB.Joins("Product").First(&offercheck,"product_id = ?", productid); err.Error != nil {
		c.JSON(500, gin.H{"Error": "No Offers"})
	} else {
		percentage := offercheck.Amount
		ProductAmount := offercheck.Product.Price
		Discount = (percentage * float64(ProductAmount)) / 100
		fmt.Println("discount:", Discount)
	}
	return Discount
}
