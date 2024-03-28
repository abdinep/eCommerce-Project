package controllers

import (
	"ecom/initializers"
	"ecom/middleware"
	"ecom/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	// "github.com/dgrijalva/jwt-go"
	// "main/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var Signup models.User
var Otp string
var Roleuser = "user"

// ============================== User Authentication =============================================
func Userlogin(c *gin.Context) {
	var form models.User
	var table models.User
	var token string
	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(501, "failed to bind json")
	}
	initializers.DB.First(&table, "email=?", strings.ToLower(form.Email))
	fmt.Println("(=======================", table, ")(====================", form.Email, "==============)")

	err = bcrypt.CompareHashAndPassword([]byte(table.Password), []byte(form.Password))
	if err != nil {
		c.JSON(501, "invalid user name or password")
	} else {
		if table.Status == "Active" {
			fmt.Println("id======>", table.ID)
			token = middleware.GenerateJwt(c, form.Email, Roleuser, table.ID)
			fmt.Println("token----->", token)
			c.SetCookie("jwtToken", token, int((time.Hour * 5).Seconds()), "/", "localhost", false, true)
			c.JSON(200, gin.H{
				"Message": "Welcome to Home page",
				"Token":token,
			})
		} else {
			c.JSON(200, "User Blocked")
		}
	}
}

//=============================== END ===============================================

func User_Logout(c *gin.Context) {

	c.SetCookie("jwtToken", "", -1, "", "", false, false)
	c.JSON(200, gin.H{"message": "Logout succesful"})
}

// ========================= Sending OTP by clicking Signup =========================
func Usersignup(c *gin.Context) {
	var check models.Otp
	er := c.ShouldBindJSON(&Signup)
	if er != nil {
		c.JSON(501, "failed to bind json")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(Signup.Password), 10)
	if err != nil {
		c.JSON(501, "Failed to hash password")
	}
	Signup.Password = string(hash)
	Otp = GenerateOtp()
	check.Otp = Otp
	err = SendOtp(Signup.Email, Otp)
	if err != nil {
		c.JSON(501, "Failed to sent otp")
	}
	result := initializers.DB.First(&check, "email=?", Signup.Email)
	if result.Error != nil {

		check = models.Otp{
			Email:     Signup.Email,
			Otp:       Otp,
			Expire_at: time.Now().Add(60 * time.Second),
		}

		initializers.DB.Create(&check)
	} else {
		initializers.DB.Model(&check).Where("email=?", Signup.Email).Updates(models.Otp{
			Otp:       Otp,
			Expire_at: time.Now().Add(60 * time.Second),
		})
	}
	// initializers.DB.Delete(&check)
	c.JSON(200, "OTP sent to your mail: "+Otp)

}

//================================== END ======================================

// ========================== OTP validation and Signup =================================
func Otpcheck(c *gin.Context) {
	var check models.Otp
	var userotp models.Otp
	var existinigOtp models.Otp
	var wallet models.Wallet
	var userid models.User
	c.ShouldBindJSON(&userotp)
	initializers.DB.First(&check, "email=?", Signup.Email)
	fmt.Println("=======(", check.Otp, ")=========(", userotp.Otp, ")=========", "(", Signup.Email, ")=========")
	value := initializers.DB.Where("otp=? AND expire_at > ?", userotp.Otp, time.Now()).First(&existinigOtp)
	if value.Error != nil {
		c.JSON(501, "Incorrect OTP or OTP expired")
	} else {
		result := initializers.DB.Create(&Signup)
		if result.Error != nil {
			c.JSON(501, "User already exist")
			return
		} else {
			initializers.DB.First(&userid, "email = ?", Signup.Email)
			wallet.Created_at = time.Now()
			wallet.UserID = userid.ID
			if err := initializers.DB.Create(&wallet); err.Error != nil {
				c.JSON(500, "Failed to create wallet")
				fmt.Println("Failed to create wallet====>", err.Error)
				return
			}
			c.JSON(200, "Successfully signed up")
		}
	}
	Signup = models.User{}
}
func Resend_Otp(c *gin.Context) {
	var check models.Otp
	Otp = GenerateOtp()
	err := SendOtp(Signup.Email, Otp)
	if err != nil {
		c.JSON(501, "Failed to sent otp")
	} else {

		result := initializers.DB.First(&check, "email=?", Signup.Email)
		if result.Error != nil {

			check = models.Otp{
				Email:     Signup.Email,
				Otp:       Otp,
				Expire_at: time.Now().Add(15 * time.Second),
			}

			result := initializers.DB.Create(&check)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, "Failed to save OTP")
			}
		} else {
			err := initializers.DB.Model(&check).Where("email=?", Signup.Email).Updates(models.Otp{
				Otp:       Otp,
				Expire_at: time.Now().Add(15 * time.Second),
			})
			if err.Error != nil {
				c.JSON(http.StatusBadRequest, "Failed to update data")
			}
		}
		c.JSON(200, "OTP sent to your mail: "+Otp)
	}

}

//============================= END =====================================

// func Login(c *gin.Context) {
// 	user := models.User{}
// 	var Find models.User
// 	c.ShouldBindJSON(&user)
// 	initializers.DB.First(&Find, "email=?", user.Email)

// 	err := bcrypt.CompareHashAndPassword([]byte(Find.Password), []byte(user.Password))
// 	if err != nil {
// 		c.JSON(200, "Invalid Username or Password.")
// 		return
// 	}
// 	if Find.Status == "Block" {
// 		c.JSON(200, "User Blocked")
// 		return
// 	}

// 	token := middleware.GenerateJwt(c, Find.Email, Roleuser, Find.ID)
// 	if err != nil {
// 		fmt.Println("TOken cant generate.")
// 	}
// 	c.SetCookie("accessToken", token, int((time.Hour * 24).Seconds()), "/", "localhost", false, true)
// 	fmt.Println(token)
// 	c.JSON(200, gin.H{
// 		"messe": "successfully Login.",
// 		"token": token,
// 	})

// }
