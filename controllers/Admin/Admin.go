package controllers

import (
	"ecom/initializers"
	"ecom/middleware"
	"ecom/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// =============================== Admin login & logout ========================
var Adminrole = "admin"
func Login(c *gin.Context) {
	var log models.Admin
	var admin models.Admin
	err := c.ShouldBindJSON(&log)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	initializers.DB.First(&admin)

	if log.Email == admin.Email && log.Password == admin.Password {
		adminID := admin.ID
		fmt.Println("==========>", admin.Email, admin.Password, adminID,"<=============")
		middleware.GenerateJwt(c, log.Email, Adminrole, adminID)
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	} else {
		c.JSON(501, gin.H{
			"message": "invalid password or Username",
		})
	}
}
func Admin_Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	middleware.BlacklistedToken[tokenString] = true
	c.JSON(200, gin.H{"message": "Logout succesful"})
}

//============================= END =======================================
