package usecase

import "test-golang/internal/aggregate"

type TodoUsecase interface {
	GetTodo() (aggregate.Todos, error)
	CreateTodo(activity string) (err error)
}
