package services

import (
	"errors"

	"github.com/bysergr/priverion_test/server/database"
	"github.com/bysergr/priverion_test/server/dto"
	"github.com/bysergr/priverion_test/server/models"
)

type UserService struct {
	userDB database.UserDB
}

func NewUserService() *UserService {
	return &UserService{userDB: database.NewUser()}
}

// Create new user
func (u *UserService) CreateUser(user models.User) error {
	return u.userDB.CreateUser(user)
}

// Get all users
func (u *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	users, err := u.userDB.GetAllUsers()
	if err != nil {
		return make([]dto.UserDTO, 0), err
	}

	return dto.UsersToDTOs(users), nil
}

// Get user by email
func (u *UserService) GetUserByEmail(email string) (dto.UserDTO, error) {
	user, err := u.userDB.GetUserByEmail(email)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return dto.UserToDTO(user), nil
}

// Get user by id
func (u *UserService) GetUserByID(id string) (dto.UserDTO, error) {
	user, err := u.userDB.GetUserByID(id)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return dto.UserToDTO(user), nil
}

// Get user by id with password
func (u *UserService) GetUserByIDPassword(id string) (dto.UserPasswordDTO, error) {
	user, err := u.userDB.GetUserByID(id)
	if err != nil {
		return dto.UserPasswordDTO{}, err
	}

	return dto.UserToPasswordDTO(user), nil
}

// Update user
func (u *UserService) UpdateUser(user dto.UserDTO, id string) error {
	oldUser, err := u.userDB.GetUserByID(id)
	if err != nil {
		return err
	}

	if oldUser.ID.Hex() == "" {
		return errors.New("user not found")
	}

	if user.Email != "" {
		oldUser.Email = user.Email
	}

	if user.Username != "" {
		oldUser.Username = user.Username
	}

	_, err = u.userDB.ChangeUser(oldUser, id)

	return err
}

// Delete user
func (u *UserService) DeleteUser(id string) error {
	return u.userDB.DeleteUser(id)
}
