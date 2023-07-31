package domain

import (
	"context"
	"errors"
)

var (
	UserNotFound      = errors.New("cannot find user")
	FailedToGetUser   = errors.New("unexpected error")
	UserAlreadyExists = errors.New("user already exists")
)

// Repository is the interface that wraps the basic CRUD operations
//
//go:generate mockery --name=Repository --output=domain --inpackage=true
type Repository interface {
	//GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByID(ctx context.Context, userID int) (*User, error)
	GetUserByEmail(ctx context.Context, userEmail string) (*User, error)
	CreateUser(ctx context.Context, name, email, password string) error
}
