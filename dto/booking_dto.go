package dto

type CreateBookingDTO struct {
	TripID uint `json:"trip_id" binding:"required"`
}
