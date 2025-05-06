package booking_handlers

import (
	"github.com/danilobml/bookatour-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(gr *gin.RouterGroup) {
	bookingGroup := gr.Group("/bookings")
	bookingGroup.Use(middlewares.Authenticate)
	bookingGroup.GET("/", getBookings)
	bookingGroup.GET("/user", getUserBookings)
	bookingGroup.GET("/tour/:id", getTourBookings)
}
