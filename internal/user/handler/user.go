package handler

import (
	"encoding/json"
	"net/http"
	"todo-list/internal/user/model"
	"todo-list/internal/user/service"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	res := service.Register(user)
	json.NewEncoder(w).Encode(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	res := service.Login(creds)
	json.NewEncoder(w).Encode(res)
}
