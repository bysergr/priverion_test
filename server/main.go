package main

import (
	"fmt"
	"net/http"

	"github.com/bysergr/priverion_test/server/utils"
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

	port := fmt.Sprintf(":%d", utils.GetENV().PORT_API)

	r.Run(port)
}
