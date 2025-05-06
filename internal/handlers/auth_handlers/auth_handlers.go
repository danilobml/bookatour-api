package auth_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/danilobml/bookatour-api/internal/models"
	"github.com/danilobml/bookatour-api/internal/services/auth_service"
	"github.com/danilobml/bookatour-api/internal/utils"
)

type User = models.User

func Signup(context *gin.Context) {
	var newUser User

	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request. " + err.Error()})
		return
	}

	newUser.Id = uuid.New().String()

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error registering user. " + err.Error()})
		return
	}
	newUser.Password = hashedPassword

	err = auth_service.RegisterUser(newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error registering user. " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered succesfully."})
}

func Login(context *gin.Context) {
	var user User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request data."})
		return
	}

	token, err := auth_service.ValidateCredentials(user.Email, user.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
