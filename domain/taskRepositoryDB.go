package domain

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbName = "todoDB"
var docCollection = "tasks"

type TaskRepositoryDb struct {
	client *mongo.Client
}

func (d TaskRepositoryDb) FindAll() ([]Task, *error) {
	db := d.client.Database(dbName).Collection(docCollection)
	cur, err := db.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var tasks []Task
	if err = cur.All(context.TODO(), &tasks); err != nil {
		log.Fatal(err)
	}
	return tasks, nil
}

func NewTaskRepositoryDb(dbClient *mongo.Client) TaskRepositoryDb {
	return TaskRepositoryDb{dbClient}
}
