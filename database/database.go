package database

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"rest-api/settings"
)

var (
	notDriverProvides = errors.New("not driver provides")
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", s.DB.User, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Name)
	switch s.DB.Engine {
	case "mariadb":
		return sqlx.ConnectContext(ctx, "mysql", connectionString)

	case "postgres":
		return sqlx.ConnectContext(ctx, "postgres", connectionString)

	default:
		return nil, notDriverProvides
	}

}
