package services

import (
	"errors"

	"github.com/bysergr/priverion_test/server/database"
	"github.com/bysergr/priverion_test/server/dto"
)

type HotelService struct{
	hotelDB database.HotelDB
}

func NewHotelService() *HotelService {
	return &HotelService{hotelDB: database.NewHotel()}
}

// Create new hotel
func (u *HotelService) CreateHotel(hotel dto.HotelDTO) error {
	newHotel, err := dto.DTOToHotel(hotel)
	if err != nil {
		return err
	}

	return u.hotelDB.CreateHotel(newHotel)
}

// Get all hotels
func (u *HotelService) GetAllHotels() ([]dto.HotelDTO, error) {
	hotels, err := u.hotelDB.GetAllHotels()
	if err != nil {
		return make([]dto.HotelDTO, 0), err
	}

	return dto.HotelsToDTOs(hotels), nil
}

// Get hotel by id
func (u *HotelService) GetHotelByID(id string) (dto.HotelDTO, error) {
	hotel, err := u.hotelDB.GetHotelByID(id)
	if err != nil {
		return dto.HotelDTO{}, err
	}

	return dto.HotelToDTO(hotel), nil
}

// Update hotel
func (u *HotelService) UpdateHotel(hotel dto.HotelDTO, id string) error {
	oldHotel, err := u.hotelDB.GetHotelByID(id)
	if err != nil {
		return err
	}

	if oldHotel.ID.Hex() == "" {
		return errors.New("hotel not found")
	}

	if hotel.Name != "" {
		oldHotel.Name = hotel.Name
	}

	if hotel.Description != "" {
		oldHotel.Description = hotel.Description
	}

	if hotel.Capacity != 0 {
		oldHotel.Capacity = hotel.Capacity
	}

	_, err = u.hotelDB.ChangeHotel(oldHotel)

	return err
}

// Delete hotel
func (u *HotelService) DeleteHotel(id string) error {
	return u.hotelDB.DeleteHotel(id)
}