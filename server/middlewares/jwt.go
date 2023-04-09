package middlewares

import (
	"fmt"
	"strings"

	"github.com/bysergr/priverion_test/server/dto"
	"github.com/bysergr/priverion_test/server/services"
	"github.com/bysergr/priverion_test/server/utils"
	"github.com/gin-gonic/gin"
)

func JWT(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	if auth == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "No token provided"})
		return
	}

	fmt.Println(strings.Replace(auth, "Bearer ", "", 1))
	token, err := utils.ValidateToken(strings.Replace(auth, "Bearer ", "", 1))
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		return
	}

	user, err := services.NewUserService().GetUserByID(token["id"])
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Invalid token"})
		return
	}

	c.Set("user", user)

	c.Next()
}

func JWTUser(c *gin.Context) {
	user := c.MustGet("user").(dto.UserDTO)

	if !(user.Role == "user" || user.Role == "admin") {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Next()
}

func JWTAdmin(c *gin.Context) {
	user := c.MustGet("user").(dto.UserDTO)

	if user.Role != "admin" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Next()
}
