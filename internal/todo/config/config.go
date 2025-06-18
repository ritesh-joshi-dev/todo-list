package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	Url        string
	Database   string
	Collection string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Url:        "mongodb://localhost:27017",
		Database:   "tododb",
		Collection: "todo",
	}
}

func ConnectDB(cfg *DBConfig) (*mongo.Database, error) {
	clientOpts := options.Client().ApplyURI(cfg.Url)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, err
	}
	return client.Database(cfg.Database), nil
}

func NewConnection() (*mongo.Database, *mongo.Collection) {
	cfg := NewDBConfig()
	db, err := ConnectDB(cfg)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}
	return db, db.Collection(cfg.Collection)
}
