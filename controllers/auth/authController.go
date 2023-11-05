package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"toxicboy/go-jwt/orm"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterController(c *gin.Context) {
	var req Register

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var userExist orm.User
	orm.DB.Where("username = ?", req.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Create new User
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	newUser := orm.User{Username: req.Username, Password: string(encryptedPassword)}
	result := orm.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
