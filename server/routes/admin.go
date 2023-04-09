package routes

import (
	"github.com/bysergr/priverion_test/server/handlers"
	"github.com/bysergr/priverion_test/server/middlewares"
	"github.com/gin-gonic/gin"
)

// AdminRouter - Routes for ADMINs
func AdminRouter(r *gin.Engine) {

	router := r.Group("/")

	// Middlewares
	router.Use(middlewares.JWTAdmin)

	// Routes for Hotel
	router.POST("/hotel", handlers.CreateHotel)
	router.PUT("/hotel/:id", handlers.UpdateHotel)
	router.DELETE("/hotel/:id", handlers.DeleteHotel)

	// Routes for Reservation
	router.GET("/reservations", handlers.GetReservations)
	router.GET("/reservation/hotel/:idHotel", handlers.GetReservationsByHotelID)
	router.PUT("/reservation/:id", handlers.UpdateReservation)
	router.DELETE("/reservation/:id", handlers.DeleteReservation)

}



