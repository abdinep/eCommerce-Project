package routers

import (
	"ecom/controllers"
	"ecom/handlers"
	"ecom/middleware"

	"github.com/gin-gonic/gin"
)

var Roleuser = "user"

func UserGroup(r *gin.RouterGroup) {
	//=========================== User sign up, Login & Logout =============================

	r.POST("/user/registration", controllers.Usersignup)
	r.POST("/user/signin", controllers.Userlogin)
	r.POST("/forgotpass/sendOTP", controllers.ForgotPassword_OTP)
	r.POST("/forgotpass/checkOTP", controllers.Forgot_Pass_OTP_Check)
	r.POST("/forgotPassword", controllers.ForgotPassword_Change)
	r.GET("/user/logout",controllers.User_Logout)

	//========================= User registration with OTP =====================================

	r.POST("/user/registration/otp", controllers.Otpcheck)
	r.POST("/user/registration/resendotp", controllers.Resend_Otp)

	//========================= User product management ========================================

	r.GET("/products",middleware.AuthMiddleware(Roleuser), handlers.ProductLoadingPage)
	r.GET("/products/details/:ID",middleware.AuthMiddleware(Roleuser), handlers.ProductDetails)
	r.POST("/products/rating",middleware.AuthMiddleware(Roleuser), handlers.RatingStrore)
	r.POST("/products/review",middleware.AuthMiddleware(Roleuser), handlers.ReviewStore)

	//========================== User Address management =======================================

	r.POST("/user/address",middleware.AuthMiddleware(Roleuser), handlers.Add_Address)
	r.GET("/user/address/:ID",middleware.AuthMiddleware(Roleuser), handlers.View_Address)
	r.PATCH("/user/address/:ID",middleware.AuthMiddleware(Roleuser), handlers.Edit_Address)
	r.DELETE("/user/address/:ID",middleware.AuthMiddleware(Roleuser), handlers.Delete_Address)

	//========================== User Cart management ==========================================

	r.POST("/cart/:ID",middleware.AuthMiddleware(Roleuser), handlers.Add_Cart)
	r.GET("/cart/:ID",middleware.AuthMiddleware(Roleuser), handlers.View_Cart)
	r.PATCH("/cart/addquantity/:ID",middleware.AuthMiddleware(Roleuser), handlers.Add_Quantity_Cart)
	r.PATCH("/cart/removequantity/:ID", handlers.Remove_Quantity_cart)
	r.DELETE("/cart/:ID",middleware.AuthMiddleware(Roleuser), handlers.Remove_Cart_Product)

	//========================== User Profile ==================================================

	r.GET("/user/profile/:ID",middleware.AuthMiddleware(Roleuser), handlers.User_Details)
	r.PATCH("user/profile/:ID",middleware.AuthMiddleware(Roleuser), handlers.Edit_Profile)
	r.GET("/user/profile/order/:ID",middleware.AuthMiddleware(Roleuser), handlers.View_Orders)
	r.DELETE("/user/profile/order/:ID",middleware.AuthMiddleware(Roleuser), handlers.Cancel_Orders)

	//============================== Checkout and Order Placing ================================

	r.POST("/checkout/:ID",middleware.AuthMiddleware(Roleuser), handlers.Checkout)

	//============================== Oauth =====================================================

	r.GET("/auth/google", controllers.Googlelogin)
	// r.GET("/auth/google/callback",controllers.GoogleCallback)

	// ================================== END ===================================================
}
