package todo

import (
	"todo-list/internal/todo/handler"

	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/api/todo/create", handler.CreateTask).Methods("POST")
	router.HandleFunc("/api/todo/get", handler.GetTask).Methods("GET")
	router.HandleFunc("/api/todo/getAll", handler.GetAllTask).Methods("GET")
	router.HandleFunc("/api/todo/update", handler.UpdateTask).Methods("POST")
	router.HandleFunc("/api/todo/completed", handler.CompletedTask).Methods("GET")
	router.HandleFunc("/api/todo/pending", handler.PendingTask).Methods("GET")
}
