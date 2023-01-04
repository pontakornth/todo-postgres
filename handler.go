package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pontakornth/todo-postgres/repository"
)

type TodoHandler struct {
	todoRepo repository.TodosRepository
}

func (handler *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	todoIdStr := chi.URLParam(r, "todoId")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID must be a number."})
		return
	}
	todo, err := handler.todoRepo.GetTodo(todoId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "No todo with given ID"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (handler *TodoHandler) GetTodoList(w http.ResponseWriter, r *http.Request) {
	todos, err := handler.todoRepo.GetAllTodos()
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(todos)
}

func (handler *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo repository.Todo
	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	createdTodo, err := handler.todoRepo.CreateTodo(newTodo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(createdTodo)
}

func (handler *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Note: Changing ID is not allowed.
	todoIdStr := chi.URLParam(r, "todoId")
	w.Header().Add("Content-Type", "application/json")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID must be a number."})
	}
	var todo repository.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid format"})
		return
	}
	todo.Id = todoId
	updatedTodo, err := handler.todoRepo.UpdateTodo(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

func (handler *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoIdStr := chi.URLParam(r, "todoId")
	w.Header().Add("Content-Type", "application/json")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID must be a number"})
		return
	}
	err = handler.todoRepo.DeleteTodo(todoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
