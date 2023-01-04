package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pontakornth/todo-postgres/repository"
)

func main() {
	r := chi.NewRouter()
	// TODO: Use configuration
	db, err := sqlx.Connect("pgx", "postgres://postgres:postgressPAssW0rd@localhost:5432/application")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewTodoRepository(db)
	handler := TodoHandler{todoRepo: repo}
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", handler.GetTodoList)
		r.Post("/", handler.CreateTodo)
		r.Route("/{todoId:^[0-9]+$}", func(r chi.Router) {
			r.Get("/", handler.GetTodo)
			r.Put("/", handler.UpdateTodo)
			r.Delete("/", handler.DeleteTodo)
		})
	})

	http.ListenAndServe(":8080", r)
}
