package service

import (
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/entity"
	"dibimbing_golang_capstone/repository"
	"strconv"
	"errors"
)

type BookingService interface {
	CreateBooking(bookingDTO dto.CreateBookingDTO, userID uint) (*entity.Booking, error)
	GetBookingsByTripID(tripID uint) ([]dto.BookingDetailDTO, error)
	GetBookingsByGuideID(guideID uint) ([]dto.BookingDetailDTO, error)
	UpdateBookingStatus(bookingID uint, guideID uint, status string) error
}

type bookingService struct {
	bookingRepository repository.BookingRepository
}

func NewBookingService(bookingRepository repository.BookingRepository) BookingService {
	return &bookingService{bookingRepository}
}

func (s *bookingService) CreateBooking(bookingDTO dto.CreateBookingDTO, userID uint) (*entity.Booking, error) {
	// Ambil trip terkait
	tripRepo := repository.NewTripRepository()
	trip, err := tripRepo.GetTripByID(strconv.Itoa(int(bookingDTO.TripID)))
	if err != nil {
		return nil, err
	}

	// Hitung jumlah booking existing
	bookings, err := s.bookingRepository.GetBookingsByTripID(bookingDTO.TripID)
	if err != nil {
		return nil, err
	}
	// Validasi: user tidak boleh booking trip yang sama dua kali
	for _, b := range bookings {
		if b.UserID == userID {
			return nil, errors.New("You have already booked this trip")
		}
	}
	if len(bookings) >= trip.Capacity {
		return nil, errors.New("Trip is fully booked")
	}

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

func (s *bookingService) GetBookingsByGuideID(guideID uint) ([]dto.BookingDetailDTO, error) {
	bookings, err := s.bookingRepository.GetBookingsByGuideID(guideID)
	if err != nil {
		return nil, err
	}
	userRepo := repository.NewUserRepository(nil)
	tripRepo := repository.NewTripRepository()
	var details []dto.BookingDetailDTO
	for _, b := range bookings {
		user, _ := userRepo.GetUserByID(b.UserID)
		trip, _ := tripRepo.GetTripByID(strconv.Itoa(int(b.TripID)))
		details = append(details, dto.BookingDetailDTO{
			BookingID:     b.BookingID,
			BookingStatus: b.BookingStatus,
			CreatedAt:     b.CreatedAt,
			UserID:        b.UserID,
			Username:      user.Username,
			Email:         user.Email,
			TripID:        b.TripID,
			City:          trip.City,
			StartDate:     trip.StartDate,
			EndDate:       trip.EndDate,
		})
	}
	return details, nil
}

func (s *bookingService) UpdateBookingStatus(bookingID uint, guideID uint, status string) error {
	// Ambil booking
	bookings, err := s.bookingRepository.GetBookingsByGuideID(guideID)
	if err != nil {
		return err
	}
	// Pastikan booking_id ada di list booking milik guide
	found := false
	for _, b := range bookings {
		if b.BookingID == bookingID {
			found = true
			break
		}
	}
	if !found {
		return errors.New("You are not authorized to update this booking")
	}
	// Update status
	return s.bookingRepository.UpdateBookingStatus(bookingID, status)
}

func (s *bookingService) GetBookingsByTripID(tripID uint) ([]dto.BookingDetailDTO, error) {
	bookings, err := s.bookingRepository.GetBookingsByTripID(tripID)
	if err != nil {
		return nil, err
	}
	// Join ke User dan Trip
	userRepo := repository.NewUserRepository(nil)
	tripRepo := repository.NewTripRepository()
	var details []dto.BookingDetailDTO
	for _, b := range bookings {
		user, _ := userRepo.GetUserByID(b.UserID)
		trip, _ := tripRepo.GetTripByID(strconv.Itoa(int(b.TripID)))
		details = append(details, dto.BookingDetailDTO{
			BookingID:     b.BookingID,
			BookingStatus: b.BookingStatus,
			CreatedAt:     b.CreatedAt,
			UserID:        b.UserID,
			Username:      user.Username,
			Email:         user.Email,
			TripID:        b.TripID,
			City:          trip.City,
			StartDate:     trip.StartDate,
			EndDate:       trip.EndDate,
		})
	}
	return details, nil
}
