package handlers

import (
	"log"
	"net/http"

	"github.com/bysergr/priverion_test/server/dto"
	ser "github.com/bysergr/priverion_test/server/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	reservationService = ser.NewReservationService()
)

// CreateReservation creates a new reservation and adds it to the database
func CreateReservation(c *gin.Context) {
	var reservation dto.ReservationDTO

	if err := c.BindJSON(&reservation); err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	reservation.ID = primitive.NewObjectID().Hex()

	idUserText := c.Param("idUser")
	idHotelText := c.Param("idHotel")

	idUser, err := primitive.ObjectIDFromHex(idUserText)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	idHotel, err := primitive.ObjectIDFromHex(idHotelText)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err = reservationService.CreateReservation(reservation, idUser, idHotel)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating reservation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation created successfully"})

}

// GetReservations returns all reservations from the database
func GetReservations(c *gin.Context) {
	reservations, err := reservationService.GetReservations()
	if err != nil {
		log.Println("Error:", err.Error())
	}

	if reservations == nil || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservations not found"})
		return
	}

	c.JSON(http.StatusOK, reservations)

}

// GetReservationByID returns a reservation by its ID
func GetReservationByID(c *gin.Context) {
	id := c.Param("id")

	reservation, err := reservationService.GetReservationByID(id)
	if err != nil {
		log.Println("Error:", err.Error())
	}

	if reservation.ID == "" || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	c.JSON(http.StatusOK, reservation)

}

// GetReservationsByUserID returns all reservations from the database by user ID
func GetReservationsByUserID(c *gin.Context) {
	id := c.Param("idUser")

	reservations, err := reservationService.GetReservationsByUserID(id)
	if err != nil {
		log.Println("Error:", err.Error())
	}

	if reservations == nil || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservations not found"})
		return
	}

	c.JSON(http.StatusOK, reservations)

}

// GetReservationsByHotelID returns all reservations from the database by hotel ID
func GetReservationsByHotelID(c *gin.Context) {
	id := c.Param("idHotel")

	reservations, err := reservationService.GetReservationsByHotelID(id)
	if err != nil {
		log.Println("Error:", err.Error())
	}

	if reservations == nil || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservations not found"})
		return
	}

	c.JSON(http.StatusOK, reservations)

}

// GetReservationsByUser returns all reservations from the database
func UpdateReservation(c *gin.Context) {
	id := c.Param("id")

	var reservation dto.ReservationDTO

	if err := c.BindJSON(&reservation); err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := reservationService.UpdateReservation(reservation, id)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating reservation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation updated successfully"})
}

// DeleteReservation deletes a reservation by its ID
func DeleteReservation(c *gin.Context) {
	id := c.Param("id")

	err := reservationService.DeleteReservation(id)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting reservation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted successfully"})
}