package tour_repository

import (
	"database/sql"

	"github.com/danilobml/bookatour-api/internal/db"
	"github.com/danilobml/bookatour-api/internal/models"
)

type Tour = models.Tour

func FindAll() ([]Tour, error) {
	query := `SELECT * FROM tours`

	rows, err := db.DB.Query(query)
	if err != nil {
		return []Tour{}, err
	}
	defer rows.Close()

	tours := []Tour{}
	for rows.Next() {
		var tour Tour
		err := rows.Scan(&tour.Id, &tour.Name, &tour.Description, &tour.Location, &tour.DateTime, &tour.UserId)
		if err != nil {
			return []Tour{}, err
		}
		tours = append(tours, tour)
	}

	return tours, nil
}

func FindById(id string) (*Tour, error) {
	query := `SELECT * FROM tours WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var tour Tour
	err := row.Scan(&tour.Id, &tour.Name, &tour.Description, &tour.Location, &tour.DateTime, &tour.UserId)
	if err != nil {
		return nil, err
	}

	return &tour, nil
}

func Save(tour Tour) (*Tour, error) {
	query := `
		INSERT INTO tours (id, name, description, location, dateTime, userId)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id, name, description, location, dateTime, userId
	`

	row := db.DB.QueryRow(query, tour.Id, tour.Name, tour.Description, tour.Location, tour.DateTime, tour.UserId)

	var saved Tour
	err := row.Scan(&saved.Id, &saved.Name, &saved.Description, &saved.Location, &saved.DateTime, &saved.UserId)
	if err != nil {
		return nil, err
	}

	return &saved, nil
}

func Update(tour Tour) (sql.Result, error) {
	query := `
		UPDATE tours 
		SET name = ?, description = ?, location = ?, dateTime = ?, userId = ?
		WHERE id = ?
	`
	return db.DB.Exec(query, tour.Name, tour.Description, tour.Location, tour.DateTime, tour.UserId, tour.Id)
}

func Delete(id string) (sql.Result, error) {
	query := `DELETE FROM tours WHERE id = ?`
	return db.DB.Exec(query, id)
}
