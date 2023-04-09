package handlers

import (
	"log"
	"net/http"

	"github.com/bysergr/priverion_test/server/dto"
	"github.com/bysergr/priverion_test/server/models"
	ser "github.com/bysergr/priverion_test/server/services"
	"github.com/bysergr/priverion_test/server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	userService = ser.NewUserService()
)

// CreateUser creates a new user and adds it to the database
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	psw, err := utils.Encrypt(user.Password)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	user.Password = psw

	user.ID = primitive.NewObjectID()

	err = userService.CreateUser(user)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// GetUsers returns all users from the database
func GetUsers(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil  {
		log.Println("Error:", err.Error())
	}

	if err != nil || users == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return 
	} 

	c.JSON(http.StatusOK, users)
}

// GetUserByID returns a user by its ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := userService.GetUserByID(id)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByEmail returns a user by its email
func GetUserByEmail(c *gin.Context) {
	var userB models.User

	if err := c.BindJSON(&userB); err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := userService.GetUserByEmail(userB.Email)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByUsername returns a user by its username
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user dto.UserDTO

	if err := c.BindJSON(&user); err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := userService.UpdateUser(user, id)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user by its ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := userService.DeleteUser(id)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
