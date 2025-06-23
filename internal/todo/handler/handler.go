package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"todo-list/internal/todo/auth"
	"todo-list/internal/todo/model"
	"todo-list/internal/todo/service"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(data, &todo); err != nil {
		log.Println("Failed to unmarshal request body:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	res := service.CreateTask(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	res := service.GetTask(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	res := service.GetAllTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	var todo model.Todo
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(data, &todo); err != nil {
		log.Println("Failed to unmarshal request body:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	res := service.UpdateTask(id, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func CompletedTask(w http.ResponseWriter, r *http.Request) {
	res := service.CompletedTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func PendingTask(w http.ResponseWriter, r *http.Request) {
	res := service.PendingTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	// In production, you'd validate a real user from DB
	userID := r.FormValue("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateJWT(userID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
