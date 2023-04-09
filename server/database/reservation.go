package database

import (
	"context"

	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type reservationDB struct{
	db *mongo.Database
}

func NewReservation () reservationDB {
	return reservationDB{db: newConnection()}
}

// Create new reservation in the database
func (u *reservationDB) CreateReservation(reservation models.Reservation) error {
	_, err := u.db.Collection("reservations").InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}

	return nil
}

// Get reservations by user id from the database
func (u *reservationDB) GetReservationsByUserID(id string) ([]models.Reservation, error) {
	var reservations []models.Reservation

	cursor, err := u.db.Collection("reservations").Find(context.TODO(), bson.M{"user": id})
	if err != nil {
		return reservations, err
	}

	err = cursor.All(context.TODO(), &reservations)
	if err != nil {
		return reservations, err
	}

	return reservations, nil
}

// Get reservations by hotel id from the database
func (u *reservationDB) GetReservationsByHotelID(id string) ([]models.Reservation, error) {
	var reservations []models.Reservation

	cursor, err := u.db.Collection("reservations").Find(context.TODO(), bson.M{"hotel": id})
	if err != nil {
		return reservations, err
	}

	err = cursor.All(context.TODO(), &reservations)
	if err != nil {
		return reservations, err
	}

	return reservations, nil
}

// Get all reservations from the database
func (u *reservationDB) GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	cursor, err := u.db.Collection("reservations").Find(context.TODO(), bson.M{})
	if err != nil {
		return reservations, err
	}

	err = cursor.All(context.TODO(), &reservations)
	if err != nil {
		return reservations, err
	}

	return reservations, nil
}

// Get reservation by id from the database
func (u *reservationDB) GetReservationByID(id string) (models.Reservation, error) {
	var reservation models.Reservation
	
	err := u.db.Collection("reservations").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&reservation)
	if err != nil {
		return reservation, err
	}

	return reservation, nil
}

// Get all reservations from the database
func (u *reservationDB) ChangeReservation(reservation models.Reservation) (models.Reservation, error) {
	var newreservation models.Reservation

	err := u.db.Collection("reservations").FindOneAndReplace(context.TODO(), bson.M{"_id": reservation.ID}, reservation).Decode(&newreservation)
	if err != nil {
		return newreservation, err
	}

	return newreservation, nil
}

// Delete reservation from the database
func (u *reservationDB) DeleteReservation(id string) error {
	_, err := u.db.Collection("reservations").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}