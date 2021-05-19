package service

import (
	"go-api/models"
	"go-api/pkg/repository"
)

type TodoItemService struct {
	itemRepository repository.TodoItem
	listRepository repository.TodoList
}

func NewTodoItemService(itemRepository repository.TodoItem, listRepository repository.TodoList) *TodoItemService {
	return &TodoItemService{
		itemRepository: itemRepository,
		listRepository: listRepository,
	}
}

func (t *TodoItemService) Create(userId, listId int, item *models.TodoItem) (int, error) {
	if _, err := t.listRepository.GetById(userId, listId); err != nil {
		return 0, err
	}
	return t.itemRepository.Create(listId, item)
}

func (t *TodoItemService) GetAllByUserList(userId, listId int) ([]models.TodoItem, error) {
	return t.itemRepository.GetAllByUserList(userId, listId)
}

func (t *TodoItemService) GetById(userId, listId, itemId int) (models.TodoItem, error) {
	return t.itemRepository.GetById(userId, listId, itemId)
}
