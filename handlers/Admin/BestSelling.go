package handlers

import (
	"ecom/initializers"
	"ecom/models"

	"github.com/gin-gonic/gin"
)

func BestSelling(c *gin.Context) {
	sortBy := c.Query("sort")

	// query := initializers.DB
	switch sortBy {
	case "product":
		var products []models.OrderItem
		query := `SELECT
		oi.product_id,
		p.product_name,
		COUNT(oi.product_id) AS total_orders
	FROM
		order_items oi
	JOIN
		products p ON oi.product_id = p.id
	GROUP BY
		oi.product_id, p.product_name
	ORDER BY
		total_orders DESC LIMIT 10;`
		initializers.DB.Raw(query).Scan(&products)

		for _, v := range products {
			c.JSON(200, gin.H{
				"Name":  v.Product.Product_Name,
				"Price": v.Product.Price,
				"ID":    v.ProductID,
			})
		}
	case "category":
		var category []models.Category
		query := `SELECT
		oi.product_id,
		p.product_name,
		COUNT(oi.product_id) AS total_orders
	FROM
		order_items oi
	JOIN
		products p ON oi.product_id = p.id
	GROUP BY
		oi.product_id, p.product_name
	ORDER BY
		total_orders DESC LIMIT 10;`
		initializers.DB.Raw(query).Scan(&category)

		// for _, v := range category {
		// 	c.JSON(200, gin.H{
		// 		"Name":  v.Product.Product_Name,
		// 		"Price": v.Product.Price,
		// 		"ID":    v.ProductID,
		// 	})
		// }
		return
	}
}
