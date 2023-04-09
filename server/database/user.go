package database

import (
	"context"

	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDB struct{
	db *mongo.Database
}

func NewUser () userDB {
	return userDB{db: newConnection()}
}

// Create new user in the database
func (u *userDB) CreateUser(user models.User) error {
	_, err := u.db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

// Get all users from the database
func (u *userDB) GetAllUsers() ([]models.User, error) {
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
func (u *userDB) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := u.db.Collection("users").FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get user by id from the database
func (u *userDB) GetUserByID(id string) (models.User, error) {
	var user models.User
	
	err := u.db.Collection("users").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get all users from the database
func (u *userDB) ChangeUser(user models.User) (models.User, error) {
	var newUser models.User

	err := u.db.Collection("users").FindOneAndReplace(context.TODO(), bson.M{"_id": user.ID}, user).Decode(&newUser)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// Delete user from the database
func (u *userDB) DeleteUser(id string) error {
	_, err := u.db.Collection("users").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}