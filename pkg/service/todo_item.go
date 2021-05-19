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

func (t *TodoItemService) Update(userId, listId, itemId int, input *models.UpdateItemListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return t.itemRepository.Update(userId, listId, itemId, input)
}

func (t *TodoItemService) DeleteById(userId, listId, itemId int) error {
	return t.itemRepository.DeleteById(userId, listId, itemId)
}
