package routers

import (
	Paymentgateways "ecom/PaymentGateways"
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
	r.GET("/user/logout", controllers.User_Logout)

	//========================= User registration with OTP =====================================

	r.POST("/user/registration/otp", controllers.Otpcheck)
	r.POST("/user/registration/resendotp", controllers.Resend_Otp)

	//========================= User product management ========================================

	r.GET("/products", middleware.JwtMiddleware(Roleuser), handlers.ProductLoadingPage)
	r.GET("/products/details/:ID", middleware.JwtMiddleware(Roleuser), handlers.ProductDetails)
	r.POST("/products/rating", middleware.JwtMiddleware(Roleuser), handlers.RatingStrore)
	r.POST("/products/review", middleware.JwtMiddleware(Roleuser), handlers.ReviewStore)

	//========================== User Address management =======================================

	r.POST("/user/address", middleware.JwtMiddleware(Roleuser), handlers.Add_Address)
	r.GET("/user/address", middleware.JwtMiddleware(Roleuser), handlers.View_Address)
	r.PATCH("/user/address/:ID", middleware.JwtMiddleware(Roleuser), handlers.Edit_Address)
	r.DELETE("/user/address/:ID", middleware.JwtMiddleware(Roleuser), handlers.Delete_Address)

	//========================== User Cart management ==========================================

	r.POST("/cart/:ID", middleware.JwtMiddleware(Roleuser), handlers.Add_Cart)
	r.GET("/cart", middleware.JwtMiddleware(Roleuser), handlers.View_Cart)
	r.PATCH("/cart/addquantity/:ID", middleware.JwtMiddleware(Roleuser), handlers.Add_Quantity_Cart)
	r.PATCH("/cart/removequantity/:ID",middleware.JwtMiddleware(Roleuser), handlers.Remove_Quantity_cart)
	r.DELETE("/cart/:ID", middleware.JwtMiddleware(Roleuser), handlers.Remove_Cart_Product)

	//========================== User Profile ==================================================

	r.GET("/user/profile", middleware.JwtMiddleware(Roleuser), handlers.User_Details)
	r.PATCH("user/profile", middleware.JwtMiddleware(Roleuser), handlers.Edit_Profile)
	r.GET("/user/profile/order", middleware.JwtMiddleware(Roleuser), handlers.View_Orders)
	r.GET("/user/profile/orderdetails/:ID", middleware.JwtMiddleware(Roleuser), handlers.View_Order_Details)
	r.PATCH("/user/profile/order/:ID", middleware.JwtMiddleware(Roleuser), handlers.Cancel_Orders)

	//============================== Checkout and Order Placing ================================

	r.POST("/checkout", middleware.JwtMiddleware(Roleuser), handlers.Checkout)

	//============================== User Advanced Search =====================================================
	r.GET("/user/search", handlers.SeaechProduct)
	//============================== Oauth =====================================================

	r.GET("/auth/google", controllers.Googlelogin)
	// r.GET("/auth/google/callback",controllers.GoogleCallback)

	//=============================== Razorpay Payment Template ========================================
	r.GET("/payment",Paymentgateways.PaymentTemplate)
	r.POST("/payment/submit",Paymentgateways.PaymentDetailsFromFrontend)

	// ================================== END ===================================================
}
