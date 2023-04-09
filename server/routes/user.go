package routes

import (
	"github.com/bysergr/priverion_test/server/handlers"
	"github.com/gin-gonic/gin"
)

// UserRouter - Routes for User login and register
func UserRouter(r *gin.Engine) {

	router := r.Group("/")

	// Routes for User
	router.PUT("/user/:id", handlers.UpdateUser)
	router.DELETE("/user/:id", handlers.DeleteUser)

	// Routes for Reservation
	router.GET("/reservation/:id", handlers.GetReservationByID)
	router.GET("/reservation/user/:idUser", handlers.GetReservationsByUserID)
	router.POST("/reservation/:idUser/:idHotel", handlers.CreateReservation)

}