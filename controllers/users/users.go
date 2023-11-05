package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toxicboy/go-jwt/orm"
)

func FindAllUser(c *gin.Context) {
	var users []orm.User
	orm.DB.Find(&users)

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
		return
	}

	c.JSON(http.StatusOK, users)
}
