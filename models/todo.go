package models

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	IsDone      string `json:"is_done" db:"is_done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (li UpdateListInput) Validate() error {
	if li.Title == nil && li.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateItemListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsDone      *bool   `json:"is_done"`
}

func (ili UpdateItemListInput) Validate() error {
	if ili.Title == nil && ili.Description == nil && ili.IsDone == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
