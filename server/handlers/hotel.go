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
	hotelService = ser.NewHotelService()
)

// CreateHotel creates a new hotel and adds it to the database
func CreateHotel(c *gin.Context) {
	var hotel dto.HotelDTO

	if err := c.BindJSON(&hotel); err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hotel.ID = primitive.NewObjectID().Hex()

	err := hotelService.CreateHotel(hotel)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel created successfully"})
}

// GetHotels returns all hotels from the database
func GetHotels(c *gin.Context) {
	hotels, err := hotelService.GetAllHotels()
	if err != nil {
		log.Println("Error:", err.Error())
	}

	if hotels == nil || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotels not found"})
		return
	}

	c.JSON(http.StatusOK, hotels)

}

// GetHotelByID returns a hotel by its ID
func GetHotelByID(c *gin.Context) {
	id := c.Param("id")

	hotel, err := hotelService.GetHotelByID(id)
	if err != nil {
		log.Println("Error:", err.Error())
	}

	if hotel.ID == "" || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	c.JSON(http.StatusOK, hotel)

}

// UpdateHotel updates a hotel by its ID
func UpdateHotel(c *gin.Context) {
	id := c.Param("id")

	var hotel dto.HotelDTO

	if err := c.BindJSON(&hotel); err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := hotelService.UpdateHotel(hotel, id)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel updated successfully"})
}

// DeleteHotel deletes a hotel by its ID
func DeleteHotel(c *gin.Context) {
	id := c.Param("id")

	err := hotelService.DeleteHotel(id)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
}
