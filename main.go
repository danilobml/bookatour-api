package main

import (
	"github.com/gin-gonic/gin"

	"github.com/danilobml/bookatour-api/internal/handlers/tours"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	tours.RegisterRoutes(api)

	router.Run(":8080")
}
