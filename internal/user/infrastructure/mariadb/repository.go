package mariadb

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"log"
)

const (
	queryInsertUser = `
         INSERT INTO users (name, email, password)
         VALUES (?, ?, ?);`
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
	return err
}
