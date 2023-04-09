package dto

import (
	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationDTO struct {
	ID        string   `json:"id" `
	Price     float64  `json:"price" `
	StartDate string   `json:"startDate" `
	EndDate   string   `json:"endDate" `
	User      UserDTO  `json:"user" `
	Hotel     HotelDTO `json:"hotel" `
}

// Convert Reservation to ReservationDTO
func ReservationToDTO(reservation models.Reservation, user models.User, hotel models.Hotel) ReservationDTO {
	return ReservationDTO{
		ID:        reservation.ID.Hex(),
		Price:     reservation.Price,
		StartDate: reservation.StartDate,
		EndDate:   reservation.EndDate,
		User:      UserToDTO(user),
		Hotel:     HotelToDTO(hotel),
	}
}

// Convert ReservationDTO to Reservation
func DTOToReservation(reservation ReservationDTO, idUser primitive.ObjectID, idHotel primitive.ObjectID) (models.Reservation, error) {
	id, err := primitive.ObjectIDFromHex(reservation.ID)
	if err != nil {
		return models.Reservation{}, err
	}

	return models.Reservation{
		ID:        id,
		Price:     reservation.Price,
		StartDate: reservation.StartDate,
		EndDate:   reservation.EndDate,
		User:      idUser,
		Hotel:     idHotel,
	}, nil
}
