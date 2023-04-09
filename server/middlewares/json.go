package middlewares

import "github.com/gin-gonic/gin"

// JSON is a middleware that sets the Content-Type header to application/json
func JSON(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
}