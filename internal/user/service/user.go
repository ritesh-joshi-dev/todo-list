package service

import (
	"todo-list/internal/todo/auth"
	"todo-list/internal/user/model"
	"todo-list/internal/user/repository"
)

func Register(user model.User) map[string]string {
	err := repository.SaveUser(user)
	if err != nil {
		return map[string]string{"status": "false", "message": err.Error()}
	}
	return map[string]string{"status": "true", "message": "User registered"}
}

func Login(creds model.User) map[string]string {
	user, err := repository.GetUser(creds.Username)
	if err != nil || user.Password != creds.Password {
		return map[string]string{"status": "false", "message": "Invalid credentials"}
	}
	token, _ := auth.GenerateJWT(user.ID)
	return map[string]string{"status": "true", "token": token}
}
