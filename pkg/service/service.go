package service

import (
	"goFirst"
	"goFirst/pkg/repository"
)

type Authorization interface {
	CreateUser(user goFirst.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list goFirst.TodoList) (int, error)
	GetAll(userId int) ([]goFirst.TodoList, error)
	GetById(userId, listId int) (goFirst.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input goFirst.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, list goFirst.TodoItem) (int, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
