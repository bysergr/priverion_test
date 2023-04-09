package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Password string             `json:"password" bson:"password,omitempty"`
	Email    string             `json:"email"`
	Role     string             `json:"role"`
}
