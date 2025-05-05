package tour_handlers

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/tours")
	group.GET("/", GetTours)
	group.GET("/:id", GetTour)
	group.POST("/", CreateNewTour)
	group.PUT("/:id", UpdateTour)
	group.DELETE("/:id", DeleteTour)
}
