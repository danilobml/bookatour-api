package main

import (
	"github.com/gin-gonic/gin"

	"github.com/danilobml/bookatour-api/internal/handlers/tour_handlers"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	tour_handlers.RegisterRoutes(api)

	router.Run(":8080")
}
