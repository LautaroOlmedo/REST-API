package application

import (
	"context"
	"errors"
	"rest-api/internal/user/domain/model"
)

var (
	InvalidParameter = errors.New("name, email and password are required")
	InvalidName      = errors.New("invalid name format")
	InvalidEmail     = errors.New("invalid email format")
	InvalidID        = errors.New("invalid ID")
	InvalidPassword  = errors.New("invalid password")
	UserAlreadyExist = errors.New("user already exists")
	UserNotFound     = errors.New("user not found")
)

// Service is the business logic of the application
//
//go:generate mockery --name=Service --output=application --inpackage
type Service interface {
	//GetAll() (map[int]*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	RegisterUser(ctx context.Context, name, email, password string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
}
