package main

import (
	"fmt"

	"github.com/bysergr/priverion_test/server/middlewares"
	"github.com/bysergr/priverion_test/server/routes"
	"github.com/bysergr/priverion_test/server/utils"
	"github.com/gin-gonic/gin"
)

// Application of gin-gonic/gin framework for Priverion
//
// Author: Sergio Rey
//
func main() {
	r := gin.Default()

	// Middlewares
	r.Use(middlewares.JSON)
	
	// Routes
	routes.BaseRouter(r)
	routes.UserRouter(r)
	routes.AdminRouter(r)


	port := fmt.Sprintf(":%d", utils.GetENV().PORT_API)

	r.Run(port)
}
