package user_repository

import (
	"github.com/danilobml/bookatour-api/internal/db"
	"github.com/danilobml/bookatour-api/internal/models"
)

type User = models.User

func FindByEmail(email string) (*User, error) {
	query := `
		SELECT *
		FROM users
		WHERE email = ?
	`

	row := db.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func Save(user User) error {
	query := `
		INSERT INTO users (id, email, password, role)
		VALUES (?, ?, ?, ?)
		RETURNING id, email, password, role
	`

	row := db.DB.QueryRow(query, user.Id, user.Email, user.Password, user.Role)

	var saved User
	err := row.Scan(&saved.Id, &saved.Email, &saved.Password, &saved.Role)
	if err != nil {
		return err
	}

	return err
}
