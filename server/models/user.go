package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty,unique"`
	Username string             `json:"username" bson:"omitempty"`
	Password string             `json:"password"`
	Email    string             `json:"email" binding:"required" unique:"true"`
	Role     string             `json:"role" binding:"required"`
}
