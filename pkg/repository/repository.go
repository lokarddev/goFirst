package repository

import (
	"github.com/jmoiron/sqlx"
	"goFirst"
)

type Authorization interface {
	CreateUser(user goFirst.User) (int, error)
	GetUser(username, password string) (goFirst.User, error)
}

type TodoList interface {
	Create(userId int, list goFirst.TodoList) (int, error)
	GetAll(userId int) ([]goFirst.TodoList, error)
	GetById(userId, listId int) (goFirst.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input goFirst.UpdateListInput) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
