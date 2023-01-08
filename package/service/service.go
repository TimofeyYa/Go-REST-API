package service

import (
	todo "todo/study"
	"todo/study/package/repository"
)

type Authorization interface {
	CreateUser(user todo.RegUser) (int, error)
	LoginUser(user todo.User) (string, error)
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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
