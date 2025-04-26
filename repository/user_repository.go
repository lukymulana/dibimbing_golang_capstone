package repository

import (
	"dibimbing_golang_capstone/config"
	"dibimbing_golang_capstone/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByUsername(username string) (*entity.User, error)
}

// instance third party
type userRepository struct{
	db *gorm.DB
} 

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
