package repository

import "github.com/jmoiron/sqlx"

type todoRepositoryPostgres struct {
	db *sqlx.DB
}

// CreateTodo implements TodosRepository
func (repo *todoRepositoryPostgres) CreateTodo(newTodo Todo) (*Todo, error) {
	var createdTodo Todo
	err := repo.db.Get(&createdTodo, "INSERT INTO todos (todo_text, is_complete) VALUES ($1, $2) RETURNING *", newTodo.TodoText, newTodo.IsComplete)
	return &createdTodo, err
}

// DeleteTodo implements TodosRepository
func (repo *todoRepositoryPostgres) DeleteTodo(id int) error {
	_, err := repo.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

// GetAllTodos implements TodosRepository
func (repo *todoRepositoryPostgres) GetAllTodos() ([]Todo, error) {
	todos := []Todo{}
	err := repo.db.Select(&todos, "SELECT * from todos ORDER BY id")
	if err != nil {
		return []Todo{}, err
	}
	return todos, nil

}

// GetTodo implements TodosRepository
func (repo *todoRepositoryPostgres) GetTodo(id int) (*Todo, error) {
	var todo Todo
	err := repo.db.Get(&todo, "SELECT * from todos WHERE id = $1", id)
	return &todo, err
}

// UpdateTodo implements TodosRepository
func (repo *todoRepositoryPostgres) UpdateTodo(newTodo Todo) (*Todo, error) {
	var updatedTodo Todo
	err := repo.db.Get(&updatedTodo, "UPDATE todos SET todo_text = $1, is_complete = $2 WHERE id = $3 RETURNING *", newTodo.TodoText, newTodo.IsComplete, newTodo.Id)
	return &updatedTodo, err
}

func NewTodoRepository(db *sqlx.DB) TodosRepository {
	return &todoRepositoryPostgres{db: db}
}
