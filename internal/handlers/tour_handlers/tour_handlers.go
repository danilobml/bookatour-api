package tour_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/danilobml/bookatour-api/internal/models/tour_models"
	"github.com/danilobml/bookatour-api/internal/services/tour_service"
)

func GetTours(context *gin.Context) {
	tours, err := tour_service.ListTours()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve from database."})
	}
	context.JSON(http.StatusOK, tours)
}

func CreateTour(context *gin.Context) {
	var newTour tour_models.Tour
	if err := context.ShouldBindJSON(&newTour); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	newTour.Id = uuid.New().String()

	saved, err := tour_service.CreateTour(newTour)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failure creating new tour."})
		return
	}

	context.JSON(http.StatusCreated, saved)
}
