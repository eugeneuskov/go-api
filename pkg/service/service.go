package service

import (
	"go-api/models"
	"go-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user *models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list *models.TodoList) (int, error)
	GetAllByUser(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	DeleteById(userId, listId int) error
	Update(userId, listId int, input *models.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item *models.TodoItem) (int, error)
	GetAllByUserList(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, listId, itemId int) (models.TodoItem, error)
	DeleteById(userId, listId, itemId int) error
	Update(userId, listId, itemId int, input *models.UpdateItemListInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repositories.Authorization),
		TodoList:      NewTodoListService(repositories.TodoList),
		TodoItem:      NewTodoItemService(repositories.TodoItem, repositories.TodoList),
	}
}
