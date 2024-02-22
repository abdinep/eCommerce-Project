package main

import (
	"ecom/controllers"
	"ecom/handlers"
	"ecom/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Dbinit()
	initializers.TableCreate()
	// controllers.GenerateOTP()
}

func main() {
	server := gin.Default()

	server.POST("/admin_login", controllers.Login)
	server.POST("/user_registration", controllers.Usersignup)
	server.POST("/user_signin", controllers.Userlogin)
	server.GET("/admin_panel/user_management", handlers.List_user)
	server.POST("/admin_panel/products/add_product", handlers.Add_Product)
	server.POST("/admin_panel/category/add_category",handlers.Category)
	server.GET("/admin_panel/category",handlers.View_Category)
	server.POST("/admin_panel/category/edit/:ID",handlers.Edit_Category)
	server.DELETE("/admin_panel/category/delete/:ID", handlers.Delete_Category)
	server.GET("/admin_panel/products", handlers.View_Product)
	server.PATCH("/admin_panel/user_management/edit/:ID", handlers.Edit_User)
	server.PATCH("/admin_panel/products/edit/:ID", handlers.Edit_Product)
	server.DELETE("/admin_panel/products/delete/:ID", handlers.Delete_Product)
	// server.PATCH("/admin_panel/products/Recover/:ID",handlers.Undelete_Product)
	server.PATCH("/admin_panel/user_management/block/:ID", handlers.Status)
	server.POST("/user_registration/otp", controllers.Otpcheck)

	server.Run(os.Getenv("PORT"))

}
