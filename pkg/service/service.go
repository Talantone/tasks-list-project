package service

import (
	tasks_app "tasks-list-project"
	"tasks-list-project/pkg/repository"
)

type Authorization interface {
	CreateUser(user tasks_app.User) (int, error)
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

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
	}
}
