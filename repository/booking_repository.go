package repository

import (
	"dibimbing_golang_capstone/config"
	"dibimbing_golang_capstone/entity"
)

type BookingRepository interface {
	CreateBooking(booking *entity.Booking) error
	GetBookingsByTripID(tripID uint) ([]entity.Booking, error)
	GetBookingsByGuideID(guideID uint) ([]entity.Booking, error)
	UpdateBookingStatus(bookingID uint, status string) error
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

// Ambil seluruh booking untuk trip yang dimiliki oleh guide tertentu
func (r *bookingRepository) GetBookingsByGuideID(guideID uint) ([]entity.Booking, error) {
	var bookings []entity.Booking
	// Join ke trip, filter trip.user_id = guideID
	if err := config.DB.Table("bookings").
		Select("bookings.*").
		Joins("JOIN trips ON bookings.trip_id = trips.trip_id").
		Where("trips.user_id = ?", guideID).
		Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

// Update status booking
func (r *bookingRepository) UpdateBookingStatus(bookingID uint, status string) error {
	return config.DB.Model(&entity.Booking{}).Where("booking_id = ?", bookingID).Update("booking_status", status).Error
}
