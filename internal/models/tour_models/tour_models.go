package tour_models

import (
	"time"
)

type Tour struct {
	Id          string    `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      string    `json:"userId"`
}
