package tour_handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/danilobml/bookatour-api/internal/models/tour_models"
	"github.com/danilobml/bookatour-api/internal/services/tour_service"
)

func GetTours(context *gin.Context) {
	tours, err := tour_service.ListTours()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving from database. " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, tours)
}

func GetTour(context *gin.Context) {
	tourId := context.Param("id")

	tour, err := tour_service.GetTourById(tourId)
	if tour == nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"message": "Tour not found."})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error finding the tour. " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, tour)
}

func CreateNewTour(context *gin.Context) {
	var newTour tour_models.Tour
	if err := context.ShouldBindJSON(&newTour); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. " + err.Error()})
		return
	}

	newTour.Id = uuid.New().String()

	saved, err := tour_service.CreateTour(newTour)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error creating new tour. " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, saved)
}

func DeleteTour(context *gin.Context) {
	tourId := context.Param("id")

	tour, err := tour_service.GetTourById(tourId)
	if tour == nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"message": "Tour not found."})
			return
		}
	}

	err = tour_service.DeleteTourById(tourId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting the tour. " + err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "Tour deleted."})
}
