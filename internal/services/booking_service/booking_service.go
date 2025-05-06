package booking_service

import (
	"errors"

	"github.com/danilobml/bookatour-api/internal/models"
	"github.com/danilobml/bookatour-api/internal/repositories/booking_repository"
)

var ErrTourNotFound = errors.New("tour not found")

type Booking = models.Booking

func GetBookingByTourIdAndUserId(tourId string, userId string) (*Booking, error) {
	return booking_repository.FindByTourIdAndUserId(tourId, userId)
}

func CancelBooking(tourId string, userId string) error {
	result, err := booking_repository.Delete(tourId, userId)
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

func ListBookings() ([]Booking, error) {
	return booking_repository.FindAll()
}

func ListUserBookings(userId string) ([]Booking, error) {
	return booking_repository.FindByUserId(userId)
}

func ListTourBookings(tourId string) ([]Booking, error) {
	return booking_repository.FindByTourId(tourId)
}
