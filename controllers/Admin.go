package controllers

import (
	"ecom/middleware"
	"ecom/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// =============================== Admin login & logout ========================
var role = "admin"

func Login(c *gin.Context) {
	var log models.Admin
	session := sessions.Default(c)
	check := session.Get("admin")
	if check != nil {
		c.JSON(200, "Already Loged in")
	} else {

		err := c.ShouldBindJSON(&log)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if log.Email == "abd@gmail.com" && log.Password == "abd@123" {
			middleware.SessionCreate(log.Email, role, c)
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
			})
		} else {
			c.JSON(501, gin.H{
				"message": "invalid password or Username",
			})
		}
	}
}
func Admin_Logout(c *gin.Context) {
	session := sessions.Default(c)
	check := session.Get("admin")
	if check == nil {
		c.JSON(200, "Not Logged in")
	} else {
		session.Delete("admin")
		session.Save()
		c.JSON(200, "Loged out")
	}
}

//============================= END =======================================
