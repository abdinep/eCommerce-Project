package routers

import (
	controllers "ecom/controllers/Admin"
	handlers "ecom/handlers/Admin"
	"ecom/middleware"

	"github.com/gin-gonic/gin"
)

var roleAdmin = "admin"

func AdminGroup(r *gin.RouterGroup) {

	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Admin_Logout)
	//=========================== Admin user management ======================================

	r.GET("/usermanagement", middleware.JwtMiddleware(roleAdmin), handlers.List_user)
	r.PATCH("/usermanagement/edit/:ID", middleware.JwtMiddleware(roleAdmin), handlers.Edit_User)
	r.PATCH("/usermanagement/block/:ID", middleware.JwtMiddleware(roleAdmin), handlers.Status)

	//=========================== Admin Coupon management ======================================

	r.POST("/coupon", handlers.Coupon)

	//=========================== Admin Product management ===================================

	r.GET("/products/addproduct", middleware.JwtMiddleware(roleAdmin), handlers.Add_Product)
	r.POST("/products/addproduct", middleware.JwtMiddleware(roleAdmin), handlers.ProductImage)
	r.GET("/products", middleware.JwtMiddleware(roleAdmin), handlers.View_Product)
	r.PATCH("/products/edit/:ID", middleware.JwtMiddleware(roleAdmin), handlers.Edit_Product)
	r.DELETE("/products/delete/:ID", middleware.JwtMiddleware(roleAdmin), handlers.Delete_Product)

	//=========================== Admin Category Management ==================================

	r.POST("/category/addcategory", middleware.JwtMiddleware(roleAdmin), handlers.Category)
	r.GET("/category", middleware.JwtMiddleware(roleAdmin), handlers.View_Category)
	r.POST("/category/edit/:ID", middleware.JwtMiddleware(roleAdmin), handlers.Edit_Category)
	r.DELETE("/category/delete/:ID", middleware.JwtMiddleware(roleAdmin), handlers.Delete_Category)
	// r.PATCH("/admin_panel/products/Recover/:ID",handlers.Undelete_Product)

	//=========================== Admin Order Management =======================================

	r.GET("/order",middleware.JwtMiddleware(roleAdmin), handlers.Admin_View_order)
	r.GET("/order/details/:ID",middleware.JwtMiddleware(roleAdmin),handlers.ViewOrderDetails)
	r.POST("/order/:ID",middleware.JwtMiddleware(roleAdmin),handlers.ChangeOrderStatus)

	//============================ Admin Offer Management ======================================

	r.POST("/offer/:ID",middleware.JwtMiddleware(roleAdmin),handlers.AddOffer)
	r.GET("/offer",middleware.JwtMiddleware(roleAdmin),handlers.ViewOffer)
}
