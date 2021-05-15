package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-api/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user *models.User) (int, error) {
	var id int

	row := r.db.QueryRow(
		fmt.Sprintf("insert into %s (name, username, password_hash) values ($1, $2, $3) returning id", usersTable),
		user.Name,
		user.Username,
		user.Password,
	)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
