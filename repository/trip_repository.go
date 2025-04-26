package repository

import (
	"dibimbing_golang_capstone/config"
	"dibimbing_golang_capstone/entity"
	"dibimbing_golang_capstone/dto"
	"errors"
)

type TripRepository interface {
	GetTripsByUserID(userID uint) ([]entity.Trip, error)
	CreateTrip(trip *entity.Trip) error
	GetTripsByCityAndDate(city, startDate, endDate string) ([]entity.Trip, error)
	GetAllTrips() ([]entity.Trip, error)
	GetTripByID(id string) (*entity.Trip, error)
	UpdateTrip(id string, tripDTO dto.CreateTripDTO, userID uint) (*entity.Trip, error)
	DeleteTrip(id string, userID uint) error
}

type tripRepository struct{}

func NewTripRepository() TripRepository {
	return &tripRepository{}
}

func (r *tripRepository) CreateTrip(trip *entity.Trip) error {
	return config.DB.Create(trip).Error
}

func (r *tripRepository) GetAllTrips() ([]entity.Trip, error) {
	var trips []entity.Trip
	if err := config.DB.Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}

func (r *tripRepository) GetTripByID(id string) (*entity.Trip, error) {
	var trip entity.Trip
	if err := config.DB.First(&trip, id).Error; err != nil {
		return nil, err
	}
	return &trip, nil
}

func (r *tripRepository) UpdateTrip(id string, tripDTO dto.CreateTripDTO, userID uint) (*entity.Trip, error) {
	var trip entity.Trip
	if err := config.DB.First(&trip, id).Error; err != nil {
		return nil, err
	}
	if trip.UserID != userID {
		return nil, errors.New("unauthorized: you can only update your own trip")
	}
	trip.City = tripDTO.City
	trip.StartDate = tripDTO.StartDate
	trip.EndDate = tripDTO.EndDate
	trip.Capacity = tripDTO.Capacity
	trip.Price = tripDTO.Price
	trip.Description = tripDTO.Description
	if err := config.DB.Save(&trip).Error; err != nil {
		return nil, err
	}
	return &trip, nil
}

func (r *tripRepository) DeleteTrip(id string, userID uint) error {
	var trip entity.Trip
	if err := config.DB.First(&trip, id).Error; err != nil {
		return err
	}
	if trip.UserID != userID {
		return errors.New("unauthorized: you can only delete your own trip")
	}
	if err := config.DB.Delete(&trip).Error; err != nil {
		return err
	}
	return nil
}

func (r *tripRepository) GetTripsByUserID(userID uint) ([]entity.Trip, error) {
	var trips []entity.Trip
	if err := config.DB.Where("user_id = ?", userID).Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}

func (r *tripRepository) GetTripsByCityAndDate(city, startDate, endDate string) ([]entity.Trip, error) {
	var trips []entity.Trip
	db := config.DB

	if city != "" {
		db = db.Where("city = ?", city)
	}
	if startDate != "" {
		db = db.Where("start_date <= ?", startDate)
	}
	if endDate != "" {
		db = db.Where("end_date >= ?", endDate)
	}

	if err := db.Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}
