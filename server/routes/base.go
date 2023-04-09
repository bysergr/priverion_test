package routes

import (
	"github.com/bysergr/priverion_test/server/handlers"
	"github.com/gin-gonic/gin"
)

// BaseRouter - Routes for not authenticated users
func BaseRouter(r *gin.Engine) {

	router := r.Group("/")

	// Routes for User
	router.GET("/users", handlers.GetUsers)
	router.GET("/user/:id", handlers.GetUserByID)
	router.GET("/user/email/", handlers.GetUserByEmail)

	// Routes for Hotel
	router.GET("/hotels", handlers.GetHotels)
	router.GET("/hotel/:id", handlers.GetHotelByID)

	// Routes for Session
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}