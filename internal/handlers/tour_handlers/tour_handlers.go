package tour_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/danilobml/bookatour-api/internal/models"
	"github.com/danilobml/bookatour-api/internal/services/tour_service"
)

type Tour = models.Tour

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
	if err != nil {
		if err == tour_service.ErrTourNotFound {
			context.JSON(http.StatusNotFound, gin.H{"message": "Tour not found."})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error finding the tour. " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, tour)
}

func CreateNewTour(context *gin.Context) {
	var newTour Tour

	err := context.ShouldBindJSON(&newTour)
	if err != nil {
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

func UpdateTour(context *gin.Context) {
	tourId := context.Param("id")

	var req models.UpdateTourRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	updatedTour := models.Tour{
		Id:          tourId,
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		DateTime:    req.DateTime,
		UserId:      req.UserId,
	}

	updated, err := tour_service.UpdateTour(updatedTour)
	if err != nil {
		if err == tour_service.ErrTourNotFound {
			context.JSON(http.StatusNotFound, gin.H{"message": "Tour not found."})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating tour: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, updated)
}

func DeleteTour(context *gin.Context) {
	tourId := context.Param("id")

	err := tour_service.DeleteTourById(tourId)
	if err != nil {
		if err == tour_service.ErrTourNotFound {
			context.JSON(http.StatusNotFound, gin.H{"message": "Tour not found."})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting tour: " + err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "Tour deleted."})
}
