package controllers

import (
	"ecom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
//=============================== Admin login ========================
func Login(c *gin.Context) {
	var log models.Admin
	err := c.ShouldBindJSON(&log)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if log.Email == "abd@gmail.com" && log.Password == "abd@123" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	} else {
		c.JSON(501, gin.H{
			"message": "invalid password or Username",
		})
	}
}
//============================= END =======================================