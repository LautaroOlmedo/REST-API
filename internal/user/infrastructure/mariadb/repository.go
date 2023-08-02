package mariadb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"rest-api/internal/user/domain"
)

const (
	queryInsertUser = `
         INSERT INTO users (name, email, password)
         VALUES (?, ?, ?);`

	queryGetUserByID = `
         SELECT id, name, email, password FROM users WHERE id = ?;`

	queryGetUSer = `
         SELECT id, name, email FROM users WHERE id = ?;`
)

type MariaDBRepository struct {
	db *sqlx.DB
}

func NewMariaDBRepository(db *sqlx.DB) *MariaDBRepository {
	return &MariaDBRepository{
		db: db,
	}
}

//func (repo *MariaDBRepository) GetAllUsers(ctx context.Context) (map[int]domain.User, error) {
//	resp, err := repo.db.QueryContext(ctx, queryGetAllUsers)
//	if err != nil {
//		return nil, domain.UserNotFound
//	}
//
//	users := make(map[int]domain.User)
//
//	for resp.Next() {
//		u := domain.User{}
//
//		err := resp.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
//		if err != nil {
//			return nil, domain.UserNotFound
//		}
//		users[u.ID] = u
//	}
//	if len(users) == 0 {
//		return nil, domain.UserNotFound
//	}
//	return users, nil
//}

func (repo *MariaDBRepository) GetUserByID(ctx context.Context, userID int) (*domain.User, error) {
	var u = domain.User{}
	err := repo.db.QueryRowContext(ctx, queryGetUSer, userID).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.UserNotFound
		}
		return nil, domain.FailedToGetUser
	}
	return &u, nil
}

func (repo *MariaDBRepository) GetUserByEmail(ctx context.Context, userEmail string) (*domain.User, error) {
	var u = domain.User{}
	err := repo.db.GetContext(ctx, &u, queryGetUSer, userEmail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.UserNotFound
		}
		return nil, domain.FailedToGetUser
	}
	return &u, nil
}

func (repo *MariaDBRepository) CreateUser(ctx context.Context, name, email, password string) error {
	_, err := repo.db.ExecContext(ctx, queryInsertUser, name, email, password)
	fmt.Println("ERROR IN MARIADB REPOSITORY", err)
	return err
}
