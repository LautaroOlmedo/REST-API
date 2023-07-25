package mariadb

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"log"
	"rest-api/internal/user/domain"
)

const (
	queryInsertUser = `
         INSERT INTO users (name, email, password)
         VALUES (?, ?, ?);
`

	queryGetAllUsers = `
         SELECT * FROM users;`
)

var (
	UserNotFound = errors.New("cannot find users")
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) SaveUser(ctx context.Context, name, email, password string) error {
	log.Printf("name: %s", name)
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
