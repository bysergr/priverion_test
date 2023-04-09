package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty,unique"`
	Name        string             `json:"name" bson:"omitempty,unique"`
	Description string             `json:"description"`
	Capacity    int                `json:"capacity"`
}
