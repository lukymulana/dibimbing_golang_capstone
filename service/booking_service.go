package service

import (
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/entity"
	"dibimbing_golang_capstone/repository"
)

type BookingService interface {
	CreateBooking(bookingDTO dto.CreateBookingDTO, userID uint) (*entity.Booking, error)
	GetBookingsByTripID(tripID uint) ([]entity.Booking, error)
}

type bookingService struct {
	bookingRepository repository.BookingRepository
}

func NewBookingService(bookingRepository repository.BookingRepository) BookingService {
	return &bookingService{bookingRepository}
}

func (s *bookingService) CreateBooking(bookingDTO dto.CreateBookingDTO, userID uint) (*entity.Booking, error) {
	booking := &entity.Booking{
		UserID:        userID,
		TripID:        bookingDTO.TripID,
		BookingStatus: "waiting",
	}

	if err := s.bookingRepository.CreateBooking(booking); err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *bookingService) GetBookingsByTripID(tripID uint) ([]entity.Booking, error) {
	return s.bookingRepository.GetBookingsByTripID(tripID)
}
