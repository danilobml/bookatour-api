package auth_service

import (
	"database/sql"
	"errors"

	"github.com/danilobml/bookatour-api/internal/models"
	"github.com/danilobml/bookatour-api/internal/repositories/user_repository"
	"github.com/danilobml/bookatour-api/internal/utils"
)

type User = models.User

var ErrUserNotFound = errors.New("tour not found")

func RegisterUser(user User) error {
	return user_repository.Save(user)
}

func ValidateCredentials(email, inputPassword string) (string, error) {
	user, err := user_repository.FindByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrUserNotFound
		}
		return "", err
	}

	isPasswordValid := utils.CheckPasswordHash(inputPassword, user.Password)
	if !isPasswordValid {
		return "", errors.New("credentials invalid")
	}

	token, err := utils.GenerateToken(user.Email, user.Id, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
