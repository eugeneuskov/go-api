package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-api/config"
)

func NewPostgresDB(config *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Db.Postgres.Host,
		config.Db.Postgres.Port,
		config.Db.Postgres.User,
		config.Db.Postgres.DbName,
		config.Db.Postgres.Password,
		config.Db.Postgres.SslMode,
	))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
