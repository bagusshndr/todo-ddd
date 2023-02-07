package aggregate

type Todos []*Todo

type Todo struct {
	ID       uint64
	Activity string
}

func NewTodo(activity string) (*Todo, error) {
	return &Todo{
		Activity: activity,
	}, nil
}

func RebuildTodos(id uint64, activity string) (*Todo, error) {
	return &Todo{
		ID:       id,
		Activity: activity,
	}, nil
}
