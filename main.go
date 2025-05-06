package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/danilobml/bookatour-api/internal/db"
	"github.com/danilobml/bookatour-api/internal/handlers/auth_handlers"
	"github.com/danilobml/bookatour-api/internal/handlers/tour_handlers"
	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	db.InitDB()

	router := gin.Default()

	api := router.Group("/api")
	auth_handlers.RegisterRoutes(api)
	tour_handlers.RegisterRoutes(api)

	router.Run(":8080")
}
