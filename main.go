package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	AuthController "toxicboy/go-jwt/controllers/auth"
	"toxicboy/go-jwt/orm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	orm.ConfigDb()
	
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/local/register", AuthController.RegisterController)

	router.Run()
}
