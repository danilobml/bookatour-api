package tour_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/danilobml/bookatour-api/internal/models"
	"github.com/danilobml/bookatour-api/internal/services/booking_service"
	"github.com/danilobml/bookatour-api/internal/services/tour_service"
)

type Tour = models.Tour
type Booking = models.Booking

func getTours(context *gin.Context) {
	tours, err := tour_service.ListTours()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving from database. " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, tours)
}

func getTour(context *gin.Context) {
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

func createNewTour(context *gin.Context) {
	var newTour Tour

	err := context.ShouldBindJSON(&newTour)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. " + err.Error()})
		return
	}

	newTour.Id = uuid.New().String()
	newTour.UserId = context.GetString("userId")

	saved, err := tour_service.CreateTour(newTour)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error creating new tour. " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, saved)
}

func updateTour(context *gin.Context) {
	tourId := context.Param("id")
	userId := context.GetString("userId")

	if !checkIsSameUser(tourId, userId, context) {
		return
	}

	var request models.UpdateTourRequest

	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	updatedTour := models.Tour{
		Id:          tourId,
		Name:        request.Name,
		Description: request.Description,
		Location:    request.Location,
		DateTime:    request.DateTime,
		UserId:      userId,
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

func deleteTour(context *gin.Context) {
	tourId := context.Param("id")
	userId := context.GetString("userId")

	if !checkIsSameUser(tourId, userId, context) {
		return
	}

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

func bookTour(context *gin.Context) {
	tourId, _ := context.Params.Get("id")
	userId := context.GetString("userId")

	if tourId == "" || userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Tour id and/or User id not supplied."})
		return
	}

	if !checkTourExists(tourId, context) {
		return
	}

	if checkBookingExists(tourId, userId) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User has already booked this tour."})
		return
	}

	bookingId := uuid.New().String()

	newBooking := Booking{
		Id: bookingId,
		TourId: tourId,
		UserId: userId,
	}

	booked, err := tour_service.BookTour(newBooking)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error: failed to book tour. " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, booked)
}

func cancelBooking(context *gin.Context) {
	tourId, _ := context.Params.Get("id")
	userId := context.GetString("userId")

	if tourId == "" || userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Tour id and/or User id not supplied."})
		return
	}

	err := booking_service.CancelBooking(tourId, userId)
	if err != nil {
		if err == tour_service.ErrTourNotFound {
			context.JSON(http.StatusNotFound, gin.H{"message": "Booking not found."})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error cancelling booking: " + err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "Booking succesfully cancelled."})
}
