package domain

import (
	"context"
)

// Repository is the interface that wraps the basic CRUD operations
type Repository interface {
	SaveUser(ctx context.Context, name, email, password string) error
	GetAllUsers(ctx context.Context) (map[int]User, error)
}
