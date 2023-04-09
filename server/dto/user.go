package dto

import (
	"github.com/bysergr/priverion_test/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserPasswordDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

// Convert User to UserPasswordDTO
func UserToPasswordDTO(user models.User) UserPasswordDTO {
	return UserPasswordDTO{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
	}
}

// Convert User to UserDTO
func UserToDTO(user models.User) UserDTO {
	return UserDTO{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}

// Convert UserDTO to User
func DTOToUser(userDTO UserDTO) models.User {
	idObject, err := primitive.ObjectIDFromHex(userDTO.ID)
	if err != nil {
		return models.User{}
	}

	return models.User{
		ID: 	 idObject,
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Role:     userDTO.Role,
	}
}

// Convert Users to DTOs
func UsersToDTOs(users []models.User) []UserDTO {
	var usersDTO []UserDTO

	for _, user := range users {
		usersDTO = append(usersDTO, UserToDTO(user))
	}

	return usersDTO
}
