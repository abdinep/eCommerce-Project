package routers

import (
	"ecom/controllers"
	"ecom/handlers"
	"ecom/middleware"

	"github.com/gin-gonic/gin"
)
var roleAdmin = "admin"
func AdminGroup(r *gin.RouterGroup) {

	r.POST("/login", controllers.Login)
	r.GET("/logout",controllers.Admin_Logout)
	//=========================== Admin user management ======================================

	r.GET("/usermanagement",middleware.AuthMiddleware(roleAdmin), handlers.List_user)
	r.PATCH("/usermanagement/edit/:ID",middleware.AuthMiddleware(roleAdmin), handlers.Edit_User)
	r.PATCH("/usermanagement/block/:ID",middleware.AuthMiddleware(roleAdmin), handlers.Status)

	//=========================== Admin Coupon management ======================================

	r.POST("/coupon", handlers.Coupon)

	//=========================== Admin Product management ===================================

	r.GET("/products/addproduct",middleware.AuthMiddleware(roleAdmin), handlers.Add_Product)
	r.POST("/products/addproduct",middleware.AuthMiddleware(roleAdmin), handlers.ProductImage)
	r.GET("/products",middleware.AuthMiddleware(roleAdmin), handlers.View_Product)
	r.PATCH("/products/edit/:ID",middleware.AuthMiddleware(roleAdmin), handlers.Edit_Product)
	r.DELETE("/products/delete/:ID",middleware.AuthMiddleware(roleAdmin), handlers.Delete_Product)

	//=========================== Admin Category Management ==================================

	r.POST("/category/addcategory",middleware.AuthMiddleware(roleAdmin), handlers.Category)
	r.GET("/category",middleware.AuthMiddleware(roleAdmin), handlers.View_Category)
	r.POST("/category/edit/:ID",middleware.AuthMiddleware(roleAdmin), handlers.Edit_Category)
	r.DELETE("/category/delete/:ID",middleware.AuthMiddleware(roleAdmin), handlers.Delete_Category)
	// r.PATCH("/admin_panel/products/Recover/:ID",handlers.Undelete_Product)

	//=========================== Admin Order Management =======================================

	r.GET("/order",middleware.AuthMiddleware(roleAdmin), handlers.Admin_View_order)
	r.POST("/order/:ID",middleware.AuthMiddleware(roleAdmin),handlers.Change_Order_Status)
}
