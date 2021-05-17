package service

import (
	"go-api/models"
	"go-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user *models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {

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
	}
}
