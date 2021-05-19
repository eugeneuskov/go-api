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
	GetById(userId, listId int) (models.TodoList, error)
	DeleteById(userId, listId int) error
	Update(userId, listId int, input *models.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item *models.TodoItem) (int, error)
	GetAllByUserList(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, listId, itemId int) (models.TodoItem, error)
	DeleteById(userId, listId, itemId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
