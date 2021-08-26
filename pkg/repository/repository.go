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
	}
}
