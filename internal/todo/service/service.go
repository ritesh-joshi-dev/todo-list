package service

import (
	"fmt"
	"todo-list/internal/todo/model"
	"todo-list/internal/todo/repository"
)

func CreateTask(todo *model.Todo) model.TodoResult {
	res, err := repository.Create(todo)
	if err != nil {
		return model.TodoResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}
	return model.TodoResult{Status: "true", Todo: *res}
}

func GetTask(id string) *model.TodoResult {
	res, err := repository.Get(id)
	if err != nil {
		return &model.TodoResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}
	return &model.TodoResult{Status: "true", Todo: *res}
}

func GetAllTask() *model.TodoAllResult {
	res, err := repository.GetAll()
	if err != nil {
		return &model.TodoAllResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}
	return &model.TodoAllResult{Status: "true", Todo: res}
}

func UpdateTask(id string, todo model.Todo) model.TodoResult {
	res, err := repository.Update(id, todo)
	if err != nil {
		return model.TodoResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}
	return model.TodoResult{Status: "true", Todo: *res}
}

func CompletedTask() *model.TodoAllResult {
	res, err := repository.GetCompletedTask()
	if err != nil {
		return &model.TodoAllResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}
	return &model.TodoAllResult{Status: "true", Todo: res}
}

func PendingTask() *model.TodoAllResult {
	res, err := repository.GetPendingTask()
	if err != nil {
		return &model.TodoAllResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}
	return &model.TodoAllResult{Status: "true", Todo: res}
}
