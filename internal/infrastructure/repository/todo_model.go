package repository

type TodoDTO struct {
	ID       uint64 `gorm:"id"`
	Activity string `gorm:"activity"`
}

func (t *TodoDTO) TableName() string {
	return "todos"
}
