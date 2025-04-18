package service

import (
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/entity"
	"dibimbing_golang_capstone/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userDTO dto.CreateUserDTO) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(userDTO dto.CreateUserDTO) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: userDTO.Username,
		Password: string(hashedPassword),
		Email:    userDTO.Email,
		Role:     userDTO.Role,
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByUsername(username string) (*entity.User, error) {
	return s.userRepository.GetUserByUsername(username)
}
