package dto

import (
	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"`
}

// Convert Hotel to HotelDTO
func HotelToDTO(hotel models.Hotel) HotelDTO {
	return HotelDTO{
		ID:          hotel.ID.Hex(),
		Name:        hotel.Name,
		Description: hotel.Description,
		Capacity:    hotel.Capacity,
	}
}

// Convert HotelDTO to Hotel
func DTOToHotel(hotelDTO HotelDTO) (models.Hotel, error) {
	idObject, err := primitive.ObjectIDFromHex(hotelDTO.ID)
	if err != nil {
		return models.Hotel{}, err
	}

	return models.Hotel{
		ID:          idObject,
		Name:        hotelDTO.Name,
		Description: hotelDTO.Description,
		Capacity:    hotelDTO.Capacity,
	}, nil
}

// Convert Hotels to DTOs
func HotelsToDTOs(hotels []models.Hotel) []HotelDTO {
	var hotelsDTO []HotelDTO

	for _, hotel := range hotels {
		hotelsDTO = append(hotelsDTO, HotelToDTO(hotel))
	}

	return hotelsDTO
}
