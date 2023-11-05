package main

import (
	"fmt"
	AuthController "toxicboy/go-jwt/controllers/auth"
	UserController "toxicboy/go-jwt/controllers/users"
	Middleware "toxicboy/go-jwt/middleware"
	"toxicboy/go-jwt/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	orm.ConfigDb()

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	app := gin.Default()
	app.Use(cors.Default())
	app.POST("local/register", AuthController.RegisterController)
	app.POST("local/login", AuthController.LoginController)

	authorized := app.Group("/users", Middleware.Auth())
	authorized.GET("", UserController.FindAllUser)

	app.Run()
}
