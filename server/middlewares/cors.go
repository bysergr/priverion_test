package middlewares

import (
	"github.com/bysergr/priverion_test/server/utils"
	"github.com/gin-gonic/gin"
)

// Cors - Middleware for CORS
func Cors(c *gin.Context) {
	env := utils.GetENV()

	c.Writer.Header().Set("Access-Control-Allow-Origin", env.CLIENT_URI)
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	c.Next()
}