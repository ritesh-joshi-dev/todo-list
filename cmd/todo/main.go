package main

import (
	"log"
	"net/http"
	"todo-list/internal/todo"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	todo.RegisterHandlers(router)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
