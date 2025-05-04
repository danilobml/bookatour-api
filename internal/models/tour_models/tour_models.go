package tour_models

import (
	"time"
)

type Tour struct {
	Id          string       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      string       `json:"userId"`
}

// TODO - remove later, for testing:
var tours = []Tour{
	{"1", "Test1", "It's a tour", "Berlin", time.Now(), "1"},
	{"2", "Test2", "It's another tour", "Hamburg", time.Now(), "1"},
}

func (tour Tour) Save() error {
	// TODO - add to database later
	tours = append(tours, tour)
	return nil
}

func GetAllTours() []Tour {
	return tours
}

