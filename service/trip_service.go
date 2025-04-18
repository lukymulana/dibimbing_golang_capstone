package service

import (
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/entity"
	"dibimbing_golang_capstone/repository"
)

type TripService interface {
	CreateTrip(tripDTO dto.CreateTripDTO, userID uint) (*entity.Trip, error)
	GetTripsByCityAndDate(city, startDate, endDate string) ([]entity.Trip, error)
}

type tripService struct {
	tripRepository repository.TripRepository
}

func NewTripService(tripRepository repository.TripRepository) TripService {
	return &tripService{tripRepository}
}

func (s *tripService) CreateTrip(tripDTO dto.CreateTripDTO, userID uint) (*entity.Trip, error) {
	trip := &entity.Trip{
		UserID:      userID,
		City:        tripDTO.City,
		StartDate:   tripDTO.StartDate,
		EndDate:     tripDTO.EndDate,
		Capacity:    tripDTO.Capacity,
		Price:       tripDTO.Price,
		Description: tripDTO.Description,
	}

	if err := s.tripRepository.CreateTrip(trip); err != nil {
		return nil, err
	}

	return trip, nil
}

func (s *tripService) GetTripsByCityAndDate(city, startDate, endDate string) ([]entity.Trip, error) {
	return s.tripRepository.GetTripsByCityAndDate(city, startDate, endDate)
}
