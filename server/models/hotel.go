package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description"`
	Capacity    int                `json:"capacity"`
}
