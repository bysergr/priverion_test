package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Application of gin-gonic/gin framework for Priverion
//
// Author: Sergio Rey
//
func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})

	r.Run()
}
