package service

import "go-api/pkg/repository"

type Authorization interface {

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
	return &Service{}
}
