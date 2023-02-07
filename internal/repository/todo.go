package repository

import "test-golang/internal/aggregate"

type TodoRepository interface {
	GetTodo() (aggregate.Todos, error)
	CreateTodo(activity string) (err error)
}
