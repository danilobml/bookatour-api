package tour_handlers

import (
	"github.com/danilobml/bookatour-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	toursGroup := rg.Group("/tours")
	toursGroup.GET("/", getTours)
	toursGroup.GET("/:id", getTour)

	authenticatedToursGroup := toursGroup.Group("/")
	authenticatedToursGroup.Use(middlewares.Authenticate)
	authenticatedToursGroup.POST("/", createNewTour)
	authenticatedToursGroup.PUT("/:id", updateTour)
	authenticatedToursGroup.DELETE("/:id", deleteTour)
	authenticatedToursGroup.POST("/:id/book", bookTour)
	authenticatedToursGroup.DELETE("/:id/cancel", cancelBooking)
}
