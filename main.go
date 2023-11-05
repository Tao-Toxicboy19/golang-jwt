package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	AuthController "toxicboy/go-jwt/controllers/auth"
	"toxicboy/go-jwt/orm"
)

func main() {
	orm.ConfigDb()

	app := gin.Default()
	app.Use(cors.Default())
	app.POST("/local/register", AuthController.RegisterController)
	app.POST("local/login", AuthController.LoginController)

	app.Run()
}
