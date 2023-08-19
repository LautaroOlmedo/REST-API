package application

import (
	"context"
	"errors"
	"rest-api/internal/user/domain/model"
)

var (
	InvalidName         = errors.New("invalid name format")
	InvalidEmail        = errors.New("invalid email format")
	InvalidID           = errors.New("invalid ID")
	InvalidPassword     = errors.New("invalid password")
	UserAlreadyExist    = errors.New("user already exists")
	UserNotFound        = errors.New("cannot find user")
	UnexpectedError     = errors.New("unexpected error")
	InternalServerError = errors.New("internal server error")
)

// Service is the business logic of the application
//
//go:generate mockery --name=Service --output=application --inpackage
type Service interface {
	GetAll(ctx context.Context) (map[int]struct {
		Name  string
		Email string
	}, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	RegisterUser(ctx context.Context, name, email, password string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
}
