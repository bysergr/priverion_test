package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reservation struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Price     float64            `json:"price" bson:"price,omitempty"`
	StartDate string             `json:"startDate" bson:"startDate"`
	EndDate   string             `json:"endDate" bson:"endDate"`
	User      primitive.ObjectID `json:"user" bson:"user,omitempty"`
	Hotel     primitive.ObjectID `json:"hotel" bson:"hotel,omitempty"`
}