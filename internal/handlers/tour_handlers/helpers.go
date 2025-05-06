package tour_handlers

import (
	"net/http"

	"github.com/danilobml/bookatour-api/internal/services/booking_service"
	"github.com/danilobml/bookatour-api/internal/services/tour_service"
	"github.com/gin-gonic/gin"
)

func checkIsSameUser(tourId string, userId string, context *gin.Context) bool {
	tour, err := tour_service.GetTourById(tourId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Tour not found."})
		return false
	}

	dbUserId := tour.UserId

	if dbUserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to execute this operation for that tour."})
		return false
	}

	return true
}

func checkTourExists(tourId string, context *gin.Context) bool {
	_, err := tour_service.GetTourById(tourId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Tour with provided id not found."})
		return false
	}

	return true
}

func checkBookingExists(tourId string, userId string) bool {
	booking, _ := booking_service.GetBookingByTourIdAndUserId(tourId, userId)
	return booking != nil
}

