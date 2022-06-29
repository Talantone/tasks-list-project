package repository

import (
	"github.com/jmoiron/sqlx"
	tasks_app "tasks-list-project"
)

type Authorization interface {
	CreateUser(user tasks_app.User) (int, error)
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
