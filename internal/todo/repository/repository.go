package repository

import (
	"context"
	"todo-list/internal/todo/config"
	"todo-list/internal/todo/model"

	"go.mongodb.org/mongo-driver/bson"
)

func Create(todo *model.Todo) (*model.Todo, error) {
	db, collection := config.NewConnection()
	_, err := collection.InsertOne(context.Background(), todo)
	_ = db.Client().Disconnect(context.Background())
	return todo, err
}

func Update(id string, todo model.Todo) (*model.Todo, error) {
	db, collection := config.NewConnection()
	filter := bson.M{"tasks.id": id}
	update := bson.M{"$set": bson.M{"tasks": todo.Tasks}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	_ = db.Client().Disconnect(context.Background())
	return &todo, err
}

func Get(id string) (*model.Todo, error) {
	db, collection := config.NewConnection()
	filter := bson.M{"tasks.id": id}
	var todo model.Todo
	err := collection.FindOne(context.Background(), filter).Decode(&todo)
	_ = db.Client().Disconnect(context.Background())
	return &todo, err
}

func GetAll() ([]*model.Todo, error) {
	db, collection := config.NewConnection()
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var todos []*model.Todo
	for cursor.Next(context.Background()) {
		var todo model.Todo
		if err := cursor.Decode(&todo); err == nil {
			todos = append(todos, &todo)
		}
	}
	_ = db.Client().Disconnect(context.Background())
	return todos, nil
}

func GetCompletedTask() ([]*model.Todo, error) {
	db, collection := config.NewConnection()
	filter := bson.M{"tasks.status": "completed"}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var todos []*model.Todo
	for cursor.Next(context.Background()) {
		var todo model.Todo
		if err := cursor.Decode(&todo); err == nil {
			todos = append(todos, &todo)
		}
	}
	_ = db.Client().Disconnect(context.Background())
	return todos, nil
}

func GetPendingTask() ([]*model.Todo, error) {
	db, collection := config.NewConnection()
	filter := bson.M{"tasks.status": "pending"}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var todos []*model.Todo
	for cursor.Next(context.Background()) {
		var todo model.Todo
		if err := cursor.Decode(&todo); err == nil {
			todos = append(todos, &todo)
		}
	}
	_ = db.Client().Disconnect(context.Background())
	return todos, nil
}
