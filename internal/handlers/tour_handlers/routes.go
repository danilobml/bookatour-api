package tour_handlers

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/tours")
	group.GET("/", GetTours)
	group.POST("/", CreateTour)
}
