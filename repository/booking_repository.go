package repository

import (
	"dibimbing_golang_capstone/config"
	"dibimbing_golang_capstone/entity"
)

type BookingRepository interface {
	CreateBooking(booking *entity.Booking) error
	GetBookingsByTripID(tripID uint) ([]entity.Booking, error)
}

type bookingRepository struct{}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{}
}

func (r *bookingRepository) CreateBooking(booking *entity.Booking) error {
	return config.DB.Create(booking).Error
}

func (r *bookingRepository) GetBookingsByTripID(tripID uint) ([]entity.Booking, error) {
	var bookings []entity.Booking
	if err := config.DB.Where("trip_id = ?", tripID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
