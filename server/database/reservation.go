package database

import (
	"context"

	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationDB struct{
	db *mongo.Database
}

func NewReservation () ReservationDB {
	return ReservationDB{db: newConnection()}
}

// Create new reservation in the database
func (u *ReservationDB) CreateReservation(reservation models.Reservation) error {
	_, err := u.db.Collection("reservations").InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}

	return nil
}

// Get reservations by user id from the database
func (u *ReservationDB) GetReservationsByUserID(id string) ([]models.Reservation, error) {
	var reservations []models.Reservation

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return reservations, err
	}

	cursor, err := u.db.Collection("reservations").Find(context.TODO(), bson.M{"user": idObject})
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
func (u *ReservationDB) GetReservationsByHotelID(id string) ([]models.Reservation, error) {
	var reservations []models.Reservation

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return reservations, err
	}

	cursor, err := u.db.Collection("reservations").Find(context.TODO(), bson.M{"hotel": idObject})
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
func (u *ReservationDB) GetAllReservations() ([]models.Reservation, error) {
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
func (u *ReservationDB) GetReservationByID(id string) (models.Reservation, error) {
	var reservation models.Reservation

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return reservation, err
	}
	
	err = u.db.Collection("reservations").FindOne(context.TODO(), bson.M{"_id": idObject}).Decode(&reservation)
	if err != nil {
		return reservation, err
	}

	return reservation, nil
}

// Get all reservations from the database
func (u *ReservationDB) ChangeReservation(reservation models.Reservation) (models.Reservation, error) {
	var newreservation models.Reservation

	err := u.db.Collection("reservations").FindOneAndReplace(context.TODO(), bson.M{"_id": reservation.ID}, reservation).Decode(&newreservation)
	if err != nil {
		return newreservation, err
	}

	return newreservation, nil
}

// Delete reservation from the database
func (u *ReservationDB) DeleteReservation(id string) error {

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = u.db.Collection("reservations").DeleteOne(context.TODO(), bson.M{"_id": idObject})
	if err != nil {
		return err
	}

	return nil
}