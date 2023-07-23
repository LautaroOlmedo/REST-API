package application

import (
	"context"
	"rest-api/internal/user/domain"
)

type RepositoryMocked struct {
}

func (rm *RepositoryMocked) SaveUser(ctx context.Context, name, email, password string) error {

	return nil
}
func (rm *RepositoryMocked) GetAllUsers(ctx context.Context) (map[int]domain.User, error) {
	return nil, nil
}
