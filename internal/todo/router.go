package todo

import (
	"net/http"
	"todo-list/internal/todo/auth"
	"todo-list/internal/todo/handler"

	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router) {
	router.Handle("/api/todo/create", auth.AuthenticateMiddleware(http.HandlerFunc(handler.CreateTask))).Methods("POST")
	router.Handle("/api/todo/get", auth.AuthenticateMiddleware(http.HandlerFunc(handler.GetTask))).Methods("GET")
	router.Handle("/api/todo/getAll", auth.AuthenticateMiddleware(http.HandlerFunc(handler.GetAllTask))).Methods("GET")
	router.Handle("/api/todo/update", auth.AuthenticateMiddleware(http.HandlerFunc(handler.UpdateTask))).Methods("POST")
	router.Handle("/api/todo/completed", auth.AuthenticateMiddleware(http.HandlerFunc(handler.CompletedTask))).Methods("GET")
	router.Handle("/api/todo/pending", auth.AuthenticateMiddleware(http.HandlerFunc(handler.PendingTask))).Methods("GET")

	// Optionally add an auth route to get token
	router.HandleFunc("/api/login", handler.GenerateToken).Methods("POST")
}
