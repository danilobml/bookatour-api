package booking_repository

import (
	"database/sql"

	"github.com/danilobml/bookatour-api/internal/db"
	"github.com/danilobml/bookatour-api/internal/models"
)

type Booking = models.Booking

func Save(booking Booking) (*Booking, error) {
	query := `
		INSERT INTO bookings (id, tourId, userId)
		VALUES (?, ?, ?)
		RETURNING id, tourId, userId
	`

	row := db.DB.QueryRow(query, booking.Id, booking.TourId, booking.UserId)

	var newBooking Booking
	err := row.Scan(&newBooking.Id, &newBooking.TourId, &newBooking.UserId)
	if err != nil {
		return nil, err
	}

	return &newBooking, nil
}

func Delete(tourId, userId string) (sql.Result, error) {
	query := `
		DELETE FROM bookings
		WHERE tourId = ? AND userId = ? 
	`
	return db.DB.Exec(query, tourId, userId)
}

func FindAll() ([]Booking, error) {
	query := `
		SELECT *
		FROM bookings
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return []Booking{}, err
	}
	defer rows.Close()

	bookings := []Booking{}
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.Id, &booking.TourId, &booking.UserId)
		if err != nil {
			return []Booking{}, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func FindAllByTourId(tourId string) ([]Booking, error) {
	query := `
		SELECT *
		FROM bookings
		WHERE tourId = ?
	`
	rows, err := db.DB.Query(query, tourId)
	if err != nil {
		return []Booking{}, err
	}
	defer rows.Close()

	bookings := []Booking{}
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.Id, &booking.TourId, &booking.UserId)
		if err != nil {
			return []Booking{}, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func FindAllByUserId(userId string) ([]Booking, error) {
	query := `
		SELECT *
		FROM bookings
		WHERE userId = ?
	`
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return []Booking{}, err
	}
	defer rows.Close()

	bookings := []Booking{}
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.Id, &booking.TourId, &booking.UserId)
		if err != nil {
			return []Booking{}, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func FindByTourIdAndUserId(tourId, userId string) (*Booking, error) {
	query := `
		SELECT *
		FROM bookings
		WHERE tourId = ? AND userId = ? 
	`
	row := db.DB.QueryRow(query, tourId, userId)

	var booking Booking
	err := row.Scan(&booking.Id, &booking.TourId, &booking.UserId)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}
