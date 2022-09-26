package tasks

import (
	"context"
	"fmt"
	"log"

	"github.com/nhanpham699/demo/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (d TaskRepositoryDb) ById(id string) (*Task, *error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	db := d.client.Database(dbName).Collection(docCollection)
	cur := db.FindOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err)
	}
	var tasks Task
	cur.Decode(&tasks)
	return &tasks, nil
}

func (d TaskRepositoryDb) Create(data dto.CreateTaskRequest) (*dto.HandleResponse, *error) {
	db := d.client.Database(dbName).Collection(docCollection)
	_, err := db.InsertOne(context.TODO(), bson.M{"name": data.Name})
	if err != nil {
		log.Fatal(err)
	}
	response := &dto.HandleResponse{
		ResultCode: "00",
		Message:    "created successfully",
	}
	return response, nil
}

func (d TaskRepositoryDb) Delete(data dto.DeleteTaskRequest) (*dto.HandleResponse, *error) {
	db := d.client.Database(dbName).Collection(docCollection)
	idPrimitive, err := primitive.ObjectIDFromHex(data.Id)
	res, err := db.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	fmt.Print(res)
	if err != nil {
		log.Fatal(err)
	}

	response := &dto.HandleResponse{
		ResultCode: "00",
		Message:    "deleted successfully",
	}

	return response, nil
}

func (d TaskRepositoryDb) Update(data dto.UpdateTaskRequest) (*dto.HandleResponse, *error) {
	db := d.client.Database(dbName).Collection(docCollection)
	idPrimitive, err := primitive.ObjectIDFromHex(data.Id)
	condition := bson.M{"_id": idPrimitive}
	update := bson.D{{"$set", bson.D{{"name", data.Name}}}}
	res, err := db.UpdateOne(context.TODO(), condition, update)
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}
	response := &dto.HandleResponse{
		ResultCode: "00",
		Message:    "updated successfully",
	}
	return response, nil
}

func NewTaskRepositoryDb(dbClient *mongo.Client) TaskRepositoryDb {
	return TaskRepositoryDb{dbClient}
}
