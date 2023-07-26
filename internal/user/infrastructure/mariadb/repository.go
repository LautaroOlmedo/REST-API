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

	queryGetAllUsers = `
         SELECT * FROM users;`

	queryGetUSer = `
         SELECT id, name, email FROM users WHERE id = ?;`
)

var (
	UserNotFound      = errors.New("cannot find user")
	FailedToGetUser   = errors.New("unexpected error")
	UserAlreadyExists = errors.New("user already exists")
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) CreateUser(ctx context.Context, name, email, password string) error {
	_, err := repo.db.ExecContext(ctx, queryInsertUser, name, email, password)
	fmt.Println(err)
	return err
}

func (repo *PostgresRepository) GetAllUsers(ctx context.Context) (map[int]domain.User, error) {
	resp, err := repo.db.QueryContext(ctx, queryGetAllUsers)
	if err != nil {
		return nil, UserNotFound
	}

	users := make(map[int]domain.User)

	for resp.Next() {
		u := domain.User{}

		err := resp.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
		if err != nil {
			return nil, UserNotFound
		}
		users[u.ID] = u
	}
	if len(users) == 0 {
		return nil, UserNotFound
	}
	return users, nil
}

func (repo *PostgresRepository) GetUserByID(ctx context.Context, userID int) (*domain.User, error) {
	var u = domain.User{}
	err := repo.db.QueryRowContext(ctx, queryGetUSer, userID).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, UserNotFound
		}
		return nil, FailedToGetUser
	}
	return &u, nil
}
