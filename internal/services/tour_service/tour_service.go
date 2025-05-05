package tour_service

import (
	"github.com/danilobml/bookatour-api/internal/models/tour_models"
	"github.com/danilobml/bookatour-api/internal/repositories/tour_repository"
)

func ListTours() ([]tour_models.Tour, error) {
	return tour_repository.FindAll()
}

func GetTourById(id string) (*tour_models.Tour, error) {
	return tour_repository.FindById(id)
}

func CreateTour(tour tour_models.Tour) (*tour_models.Tour, error) {
	savedTour, err := tour_repository.Save(tour)
	return savedTour, err
}

func UpdateTour(tour tour_models.Tour) (*tour_models.Tour, error) {
	updatedTour, err := tour_repository.Update(tour)
	return updatedTour, err
}

func DeleteTourById(id string) error {
	err := tour_repository.Delete(id)
	return err
}
