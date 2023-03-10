package repository

import (
	todo "todo/study"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.RegUser) (int, error)
	IsUserExist(user todo.User) (int, error)
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
