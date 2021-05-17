package service

import (
	"go-api/models"
	"go-api/pkg/repository"
)

type TodoListService struct {
	repository repository.TodoList
}

func NewTodoListService(repository repository.TodoList) *TodoListService {
	return &TodoListService{repository: repository}
}

func (s *TodoListService) Create(userId int, list *models.TodoList) (int, error) {
	return s.repository.Create(userId, list)
}
