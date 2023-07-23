package application

import (
	"context"
	"errors"
	"net/mail"
	"rest-api/internal/user/domain"
)

var (
	InvalidParameter = errors.New("name, email and password are required")
	InvalidEmail     = errors.New("invalid email format")
)

type UserService struct {
	userRepository domain.Repository
}

func NewUserService(userRepository domain.Repository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) SaveUser(ctx context.Context, name, email, password string) error {
	if name == "" || email == "" || password == "" {
		return InvalidParameter
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return InvalidEmail
	}
	return s.userRepository.SaveUser(ctx, name, email, password)
}

func (s *UserService) GetAllUsers(ctx context.Context) (map[int]domain.User, error) {
	return s.userRepository.GetAllUsers(ctx)
}
