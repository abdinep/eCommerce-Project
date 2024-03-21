package main

import (
	"ecom/initializers"
	"ecom/routers"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	server.LoadHTMLGlob("templates/*")
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))

	user := server.Group("/")
	routers.UserGroup(user)

	admin := server.Group("/admin")
	routers.AdminGroup(admin)

	server.Run(os.Getenv("PORT"))

}
