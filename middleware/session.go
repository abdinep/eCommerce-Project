package middleware

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionCreate(email string, role string, c *gin.Context) {
	session := sessions.Default(c)
	session.Set(role, email)
	err := session.Save()
	check := session.Get(role)
	fmt.Println("========+++", check)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create session",
		})
	} else {
		return
	}
}
func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		check := session.Get(role)
		fmt.Println("==========---",check)
		if check == nil{
			c.JSON(200,gin.H{
				"message": "Unautherised",
			})
			c.Abort()
		}
		c.Next()
	}
}
