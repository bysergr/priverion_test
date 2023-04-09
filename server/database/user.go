package database

import (
	"context"

	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDB struct{
	db *mongo.Database
}

func NewUser () UserDB {
	return UserDB{db: newConnection()}
}

// Create new user in the database
func (u *UserDB) CreateUser(user models.User) error {
	_, err := u.db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

// Get all users from the database
func (u *UserDB) GetAllUsers() ([]models.User, error) {
	var users []models.User

	cursor, err := u.db.Collection("users").Find(context.TODO(), bson.M{})
	if err != nil {
		return users, err
	}

	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return users, err
	}

	return users, nil
}

// Get user by email from the database
func (u *UserDB) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := u.db.Collection("users").FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get user by id from the database
func (u *UserDB) GetUserByID(id string) (models.User, error) {
	var user models.User

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	
	err = u.db.Collection("users").FindOne(context.TODO(), bson.M{"_id": idObject}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get all users from the database
func (u *UserDB) ChangeUser(user models.User, id string) (models.User, error) {
	var newUser models.User

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	

	err = u.db.Collection("users").FindOneAndReplace(context.TODO(), bson.M{"_id": idObject}, user).Decode(&newUser)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// Delete user from the database
func (u *UserDB) DeleteUser(id string) error {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = u.db.Collection("users").DeleteOne(context.TODO(), bson.M{"_id": idObject})
	if err != nil {
		return err
	}

	return nil
}