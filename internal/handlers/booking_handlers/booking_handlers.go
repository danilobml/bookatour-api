package booking_handlers

import (
	"net/http"

	"github.com/danilobml/bookatour-api/internal/services/booking_service"
	"github.com/danilobml/bookatour-api/internal/handlers/tour_handlers"
	"github.com/gin-gonic/gin"
)

func getBookings(context *gin.Context) {
	bookings, err := booking_service.ListBookings()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get bookings. " + err.Error()})
	}

	context.JSON(http.StatusOK, bookings)
}

func getUserBookings(context *gin.Context) {
	userId := context.GetString("userId")

	bookings, err := booking_service.ListUserBookings(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get bookings. " + err.Error()})
	}

	context.JSON(http.StatusOK, bookings)
}

func getTourBookings(context *gin.Context) {
	tourId, _ := context.Params.Get("id")
	if tourId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get Tour id."})
	}

	if !tour_handlers.CheckTourExists(tourId, context) {
		return
	}

	bookings, err := booking_service.ListTourBookings(tourId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get bookings. " + err.Error()})
	}

	context.JSON(http.StatusOK, bookings)
}
