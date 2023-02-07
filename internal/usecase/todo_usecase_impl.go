package usecase

import (
	"test-golang/internal/aggregate"
	"test-golang/internal/repository"
)

type TodoUsecaseImpl struct {
	repo repository.TodoRepository
}

func (t *TodoUsecaseImpl) GetTodo() (aggregate.Todos, error) {
	todo, err := t.repo.GetTodo()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *TodoUsecaseImpl) CreateTodo(activity string) (err error) {
	err = t.repo.CreateTodo(activity)
	if err != nil {
		return err
	}
	return nil
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &TodoUsecaseImpl{
		repo: repo,
	}
}
