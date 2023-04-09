package handlers

import (
	"log"
	"net/http"

	"github.com/bysergr/priverion_test/server/dto"
	"github.com/bysergr/priverion_test/server/models"
	"github.com/bysergr/priverion_test/server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// CreateUser creates a new user and adds it to the database
func Register(c *gin.Context) {
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

	auth := dto.Auth{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Password: user.Password,
	}

	token, exp, err := utils.GenerateToken(auth)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Generating Token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "id": user.ID, "type": "Bearer", "exp": exp, "user": user.Email})
}

func Login(c *gin.Context) {
	var credentials dto.Auth

	if err := c.BindJSON(&credentials); err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := userService.GetUserByIDPassword(credentials.ID)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !utils.CompareEncrypt(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	auth := dto.Auth{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}

	token, exp, err := utils.GenerateToken(auth)
	if err != nil {
		log.Println("Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Generating Token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "id": user.ID, "type": "Bearer", "exp": exp, "user": user.Email})
}