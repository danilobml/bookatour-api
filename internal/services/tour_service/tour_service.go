package tour_service

import (
	"database/sql"
	"errors"

	"github.com/danilobml/bookatour-api/internal/models"
	"github.com/danilobml/bookatour-api/internal/repositories/tour_repository"
)

var ErrTourNotFound = errors.New("tour not found")

type Tour = models.Tour

func ListTours() ([]Tour, error) {
	return tour_repository.FindAll()
}

func GetTourById(id string) (*Tour, error) {
	tour, err := tour_repository.FindById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTourNotFound
		}
		return nil, err
	}
	return tour, nil
}

func CreateTour(tour Tour) (*Tour, error) {
	return tour_repository.Save(tour)
}

func UpdateTour(tour Tour) (*Tour, error) {
	result, err := tour_repository.Update(tour)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, ErrTourNotFound
	}
	return &tour, nil
}

func DeleteTourById(id string) error {
	result, err := tour_repository.Delete(id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrTourNotFound
	}
	return nil
}
