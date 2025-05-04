package tour_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/danilobml/bookatour-api/internal/models/tour_models"
)

func GetTours(context *gin.Context) {
	tours := tour_models.GetAllTours()
	context.JSON(http.StatusOK, tours)
}

func CreateTour(context *gin.Context) {
	var newTour tour_models.Tour
	err := context.ShouldBindJSON(&newTour)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	newTour.Id = uuid.New().String()
	
	err = newTour.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failure creating new tour."})
		return
	}

	context.JSON(http.StatusCreated, newTour)
}