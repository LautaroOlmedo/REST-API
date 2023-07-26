package application

import (
	"context"
	"errors"
	"net/mail"
	"rest-api/internal/user/domain"
)

var (
	InvalidParameter  = errors.New("name, email and password are required")
	InvalidEmail      = errors.New("invalid email format")
	InvalidID         = errors.New("invalid ID")
	UserAlrreadyExist = errors.New("user already exist")
)

type UserService struct {
	userRepository domain.Repository
}

func NewUserService(userRepository domain.Repository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(ctx context.Context, name, email, password string) error {
	if name == "" || email == "" || password == "" {
		return InvalidParameter
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return InvalidEmail
	}
	return s.userRepository.CreateUser(ctx, name, email, password)
}

func (s *UserService) GetAllUsers(ctx context.Context) (map[int]domain.User, error) {
	return s.userRepository.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, userID int) (*domain.User, error) {
	if userID < 1 {
		return nil, InvalidID
	}
	return s.userRepository.GetUserByID(ctx, userID)
}
