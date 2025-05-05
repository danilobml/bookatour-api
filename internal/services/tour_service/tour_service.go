package tour_service

import (
	"github.com/danilobml/bookatour-api/internal/models/tour_models"
	"github.com/danilobml/bookatour-api/internal/repositories/tour_repository"
)

func ListTours() ([]tour_models.Tour, error) {
	return tour_repository.GetAll()
}

func CreateTour(tour tour_models.Tour) (tour_models.Tour, error) {
	err := tour_repository.Save(tour)
	return tour, err
}