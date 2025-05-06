package models

type Booking struct {
	Id string `json:"id"`
	TourId string `json:"tourId" binding:"required"`
	UserId string `json:"userId" binding:"required"`
}
