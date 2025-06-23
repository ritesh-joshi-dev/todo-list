package repository

import (
	"errors"
	"todo-list/internal/user/model"
)

var users = make(map[string]model.User) // temp in-memory storage; replace with MongoDB

func SaveUser(user model.User) error {
	if _, exists := users[user.Username]; exists {
		return errors.New("user already exists")
	}
	users[user.Username] = user
	return nil
}

func GetUser(username string) (model.User, error) {
	user, exists := users[username]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}
