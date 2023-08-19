package domain

import (
	"context"
)

// Repository is the interface that wraps the basic CRUD operations
//
//go:generate mockery --name=Repository --output=domain --inpackage=true
type Repository interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, userID int) (*User, error)
	GetUserByEmail(ctx context.Context, userEmail string) (*User, error)
	CreateUser(ctx context.Context, name, email, password string) error
}
