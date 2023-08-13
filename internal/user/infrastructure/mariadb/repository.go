package mariadb

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"rest-api/internal/user/application"
	"rest-api/internal/user/domain"
)

const (
	queryInsertUser = `
         INSERT INTO users (name, email, password)
         VALUES (?, ?, ?);`

	queryGetUserByEmail = `
         SELECT id, name, email, password FROM users WHERE email = ?;`

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
	var u = &domain.User{}
	err := repo.db.QueryRowContext(ctx, queryGetUSer, userID).Scan(u.ID, u.Name, u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, application.UserNotFound
		}
		return nil, application.InternalServerError
	}
	return u, nil
}

func (repo *MariaDBRepository) GetUserByEmail(ctx context.Context, userEmail string) (*domain.User, error) {
	var u = domain.User{}
	err := repo.db.GetContext(ctx, &u, queryGetUserByEmail, userEmail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, application.UserNotFound
		}
		return nil, application.InternalServerError
	}

	return &u, nil
}

func (repo *MariaDBRepository) CreateUser(ctx context.Context, name, email, password string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return application.InternalServerError
	}
	_, err = tx.ExecContext(ctx, queryInsertUser, name, email, password)
	if err != nil {
		tx.Rollback()
		return application.InternalServerError
	}

	err = tx.Commit()
	if err != nil {
		return application.InternalServerError
	}
	return nil
}
