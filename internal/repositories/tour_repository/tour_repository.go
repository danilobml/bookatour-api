package tour_repository

import (
	"github.com/danilobml/bookatour-api/internal/db"
	"github.com/danilobml/bookatour-api/internal/models/tour_models"
)

func FindAll() ([]tour_models.Tour, error) {
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

func FindById(id string) (*tour_models.Tour, error) {
	query := `
		SELECT *
		FROM tours
		WHERE id = ?
	`
	
	row := db.DB.QueryRow(query, id)

	var tour tour_models.Tour
	err := row.Scan(&tour.Id, &tour.Name, &tour.Description, &tour.Location, &tour.DateTime, &tour.UserId)
	if err != nil {
		return nil, err
	}

	return &tour, nil
}

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

func Delete(id string) error {
	query := `
		DELETE FROM tours
		WHERE id = ?
	`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
