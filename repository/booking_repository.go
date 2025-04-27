package repository

import (
	"dibimbing_golang_capstone/entity"
	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *entity.Booking) error
	GetBookingsByTripID(tripID uint) ([]entity.Booking, error)
	GetBookingsByGuideID(guideID uint) ([]entity.Booking, error)
	UpdateBookingStatus(bookingID uint, status string) error
	DeleteBooking(bookingID uint) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking *entity.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingsByTripID(tripID uint) ([]entity.Booking, error) {
	var bookings []entity.Booking
	if err := r.db.Where("trip_id = ?", tripID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) GetBookingsByGuideID(guideID uint) ([]entity.Booking, error) {
	var bookings []entity.Booking
	if err := r.db.Table("bookings").
		Select("bookings.*").
		Joins("JOIN trips ON bookings.trip_id = trips.trip_id").
		Where("trips.user_id = ?", guideID).
		Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) UpdateBookingStatus(bookingID uint, status string) error {
	return r.db.Model(&entity.Booking{}).Where("booking_id = ?", bookingID).Update("booking_status", status).Error
}

func (r *bookingRepository) DeleteBooking(bookingID uint) error {
	return r.db.Delete(&entity.Booking{}, bookingID).Error
}
