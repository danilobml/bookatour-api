package tour_repository

import (
	"github.com/danilobml/bookatour-api/internal/db"
	"github.com/danilobml/bookatour-api/internal/models/tour_models"
)

func Save(tour tour_models.Tour) error {
	query := `
		INSERT INTO tours(id, name, description, location, dateTime, userId)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(tour.Id, tour.Name, tour.Description, tour.Location, tour.DateTime, tour.UserId)
	if err != nil {
		return err
	}
	return nil
}

func GetAll() ([]tour_models.Tour, error) {
	query := `SELECT * FROM tours`

	rows, err := db.DB.Query(query)
	if err != nil {
		return []tour_models.Tour{}, err
	}
	defer rows.Close()

	tours := []tour_models.Tour{}
	for rows.Next() {
		var tour tour_models.Tour

		err := rows.Scan(&tour.Id, &tour.Name, &tour.Description, &tour.Location, &tour.DateTime, &tour.UserId)
		if err != nil {
			return []tour_models.Tour{}, err
		}

		tours = append(tours, tour)
	}

	return tours, nil
}
