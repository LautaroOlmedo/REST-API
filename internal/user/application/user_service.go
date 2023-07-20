package application

import (
	"context"
	"rest-api/internal/user/domain"
)

type UserService struct {
	userRepository domain.Repository
}

func NewUserService(userRepository domain.Repository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) SaveUser(ctx context.Context, name, email, password string) error {
	return s.userRepository.SaveUser(ctx, name, email, password)
}
