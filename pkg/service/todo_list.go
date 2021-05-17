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

func (s *TodoListService) GetAllByUser(userId int) ([]models.TodoList, error) {
	return s.repository.GetAllByUser(userId)
}

func (s *TodoListService) GetById(userId, listId int) (models.TodoList, error) {
	return s.repository.GetById(userId, listId)
}

func (s *TodoListService) DeleteById(userId, listId int) error {
	return s.repository.DeleteById(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input *models.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repository.Update(userId, listId, input)
}
