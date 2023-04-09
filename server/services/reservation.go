package services

import (
	"errors"

	"github.com/bysergr/priverion_test/server/database"
	"github.com/bysergr/priverion_test/server/dto"
	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	reservationDB database.ReservationDB
}

func NewReservationService() *ReservationService {
	return &ReservationService{reservationDB: database.NewReservation()}
}

// Create new reservation
func (u *ReservationService) CreateReservation(reservation dto.ReservationDTO, idUser primitive.ObjectID, idHotel primitive.ObjectID) error {
	newReservation, err := dto.DTOToReservation(reservation, idUser, idHotel)
	if err != nil {
		return err
	}

	err = u.reservationDB.CreateReservation(newReservation)
	if err != nil {
		return err
	}

	return nil
}

// Get all reservations
func (u *ReservationService) GetReservations() ([]dto.ReservationDTO, error) {
	reservations, err := u.reservationDB.GetAllReservations()
	if err != nil {
		return nil, err
	}

	reservationsDTO, err := ReservationsToDTOs(reservations)
	if err != nil {
		return nil, err
	}

	return reservationsDTO, nil
}

// Get reservations by id
func (u *ReservationService) GetReservationByID(id string) (dto.ReservationDTO, error) {
	reservation, err := u.reservationDB.GetReservationByID(id)
	if err != nil {
		return dto.ReservationDTO{}, err
	}

	reservations := []models.Reservation{reservation}

	reservationsDTO, err := ReservationsToDTOs(reservations)
	if err != nil {
		return dto.ReservationDTO{}, err
	}

	return reservationsDTO[0], nil
}

// Get reservations by hotel id
func (u *ReservationService) GetReservationsByHotelID(id string) ([]dto.ReservationDTO, error) {
	reservations, err := u.reservationDB.GetReservationsByHotelID(id)
	if err != nil {
		return []dto.ReservationDTO{}, err
	}

	reservationsDTO, err := ReservationsToDTOs(reservations)
	if err != nil {
		return []dto.ReservationDTO{}, err
	}

	return reservationsDTO, nil
}

// Get reservations by user id
func (u *ReservationService) GetReservationsByUserID(id string) ([]dto.ReservationDTO, error) {
	reservations, err := u.reservationDB.GetReservationsByUserID(id)
	if err != nil {
		return []dto.ReservationDTO{}, err
	}

	reservationsDTO, err := ReservationsToDTOs(reservations)
	if err != nil {
		return []dto.ReservationDTO{}, err
	}

	return reservationsDTO, nil
}

// Update reservation
func (u *ReservationService) UpdateReservation(reservation dto.ReservationDTO, id string) error {
	oldReservation, err := u.reservationDB.GetReservationByID(id)
	if err != nil {
		return err
	}

	if oldReservation.ID.Hex() == "" {
		return errors.New("reservation not found")
	}

	if reservation.StartDate != "" {
		oldReservation.StartDate = reservation.StartDate
	}

	if reservation.EndDate != "" {
		oldReservation.EndDate = reservation.EndDate
	}

	if reservation.Price != 0 {
		oldReservation.Price = reservation.Price
	}

	_, err = u.reservationDB.ChangeReservation(oldReservation)

	return err
}

// Delete reservation
func (u *ReservationService) DeleteReservation(id string) error {
	return u.reservationDB.DeleteReservation(id)
}


// Convert reservations to reservationsDTO
func ReservationsToDTOs(reservations []models.Reservation) ([]dto.ReservationDTO, error) {

	var reservationsDTO []dto.ReservationDTO

	userService := NewUserService()
	hotelService := NewHotelService()

	for _, reservation := range reservations {

		user, err := userService.GetUserByID(reservation.User.Hex())
		if err != nil {
			return nil, errors.New("not could make reservation because not found user")
		}

		hotel, err := hotelService.GetHotelByID(reservation.Hotel.Hex())
		if err != nil {
			return nil, errors.New("not could make reservation because not found hotel")
		}

		eHotel, err := dto.DTOToHotel(hotel)
		if err != nil {
			return nil, err
		}

		reservationsDTO = append(reservationsDTO, dto.ReservationToDTO(reservation, dto.DTOToUser(user), eHotel))
	}

	return reservationsDTO, nil
}
