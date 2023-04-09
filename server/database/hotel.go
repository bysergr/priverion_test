package database

import (
	"context"

	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelDB struct {
	db *mongo.Database
}

func NewHotel() HotelDB {
	return HotelDB{db: newConnection()}
}

// Create new hotel in the database
func (u *HotelDB) CreateHotel(hotel models.Hotel) error {
	_, err := u.db.Collection("hotels").InsertOne(context.TODO(), hotel)
	if err != nil {
		return err
	}

	return nil
}

// Get all hotels from the database
func (u *HotelDB) GetAllHotels() ([]models.Hotel, error) {
	var hotels []models.Hotel

	cursor, err := u.db.Collection("hotels").Find(context.TODO(), bson.M{})
	if err != nil {
		return hotels, err
	}

	err = cursor.All(context.TODO(), &hotels)
	if err != nil {
		return hotels, err
	}

	return hotels, nil
}

// Get hotel by id from the database
func (u *HotelDB) GetHotelByID(id string) (models.Hotel, error) {
	var hotel models.Hotel

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return hotel, err
	}

	err = u.db.Collection("hotels").FindOne(context.TODO(), bson.M{"_id": idObject}).Decode(&hotel)
	if err != nil {
		return hotel, err
	}

	return hotel, nil
}

// Get all hotels from the database
func (u *HotelDB) ChangeHotel(hotel models.Hotel) (models.Hotel, error) {
	var newhotel models.Hotel

	err := u.db.Collection("hotels").FindOneAndReplace(context.TODO(), bson.M{"_id": hotel.ID}, hotel).Decode(&newhotel)
	if err != nil {
		return newhotel, err
	}

	return newhotel, nil
}

// Delete hotel from the database
func (u *HotelDB) DeleteHotel(id string) error {

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = u.db.Collection("hotels").DeleteOne(context.TODO(), bson.M{"_id": idObject})
	if err != nil {
		return err
	}

	return nil
}
