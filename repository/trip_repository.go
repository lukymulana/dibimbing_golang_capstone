package repository

import (
	"dibimbing_golang_capstone/config"
	"dibimbing_golang_capstone/entity"
)

type TripRepository interface {
	CreateTrip(trip *entity.Trip) error
	GetTripsByCityAndDate(city, startDate, endDate string) ([]entity.Trip, error)
}

type tripRepository struct{}

func NewTripRepository() TripRepository {
	return &tripRepository{}
}

func (r *tripRepository) CreateTrip(trip *entity.Trip) error {
	return config.DB.Create(trip).Error
}

func (r *tripRepository) GetTripsByCityAndDate(city, startDate, endDate string) ([]entity.Trip, error) {
	var trips []entity.Trip
	if err := config.DB.Where("city = ? AND start_date <= ? AND end_date >= ?", city, startDate, endDate).Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}
