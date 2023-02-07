package repository

import (
	"database/sql"
	"errors"
	"test-golang/internal/aggregate"
	"test-golang/internal/repository"
)

type TodoRepositoryMysql struct {
	db *sql.DB
}

func NewTodoRepository(Conn *sql.DB) repository.TodoRepository {
	return &TodoRepositoryMysql{
		db: Conn,
	}
}

func (t *TodoRepositoryMysql) fetch(query string, args ...interface{}) (aggregate.Todos, error) {
	rows, err := t.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	todoDTOs := []TodoDTO{}

	for rows.Next() {
		t := TodoDTO{}
		err = rows.Scan(
			&t.ID,
			&t.Activity,
		)

		if err != nil {
			return nil, err
		}

		todoDTOs = append(todoDTOs, t)
	}

	todos := aggregate.Todos{}
	for _, todoDTO := range todoDTOs {
		aggregateTodos, _ := aggregate.RebuildTodos(todoDTO.ID, todoDTO.Activity)

		todos = append(todos, aggregateTodos)
	}

	return todos, nil
}

func (t *TodoRepositoryMysql) GetTodo() (res aggregate.Todos, err error) {
	query := `SELECT id, activity FROM todos`

	res, err = t.fetch(query)
	if err != nil {
		return nil, errors.New("todo not found")
	}
	return
}

func (t *TodoRepositoryMysql) CreateTodo(activity string) (err error) {
	query := "INSERT INTO todos (activity) VALUES (?)"
	_, err = t.db.Exec(query, activity)
	if err != nil {
		return err
	}
	return
}
