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
}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repositories.Authorization),
		TodoList: NewTodoListService(repositories.TodoList),
	}
}
