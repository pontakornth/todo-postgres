package repository

type Todo struct {
	Id         int    `db:"id" json:"id"`
	TodoText   string `db:"todo_text" json:"todoText"`
	IsComplete bool   `db:"is_complete" json:"isComplete"`
}

type TodosRepository interface {
	GetAllTodos() ([]Todo, error)
	GetTodo(id int) (*Todo, error)
	UpdateTodo(newTodo Todo) (*Todo, error)
	CreateTodo(newTodo Todo) (*Todo, error)
	DeleteTodo(id int) error
}

func NewTodo(text string, isComplete bool) Todo {
	return Todo{Id: 0, TodoText: text, IsComplete: isComplete}
}
