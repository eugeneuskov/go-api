package repository

import (
	"github.com/jmoiron/sqlx"
	"go-api/models"
)

type Authorization interface {
	CreateUser(user *models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list *models.TodoList) (int, error)
	GetAllByUser(userId int) ([]models.TodoList, error)
}

type TodoItem interface {

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
	}
}
