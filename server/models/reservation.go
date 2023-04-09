package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reservation struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty,unique"`
	Price     float64            `json:"price" bson:"omitempty"`
	StartDate string             `json:"startDate" bson:"omitempty"`
	EndDate   string             `json:"endDate" bson:"omitempty"`
	User      primitive.ObjectID `json:"user" bson:"omitempty"`
	Hotel     primitive.ObjectID `json:"hotel" bson:"omitempty"`
}