package driver

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

var user = "nhanpham"
var password = "nhan123"
var mongoUri = "mongodb+srv://" + user + ":" + password + "@cluster0.c3e8c.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func ConnectMongoDB() *MongoDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("connnected to MongoDB!")
	Mongo.Client = client
	return Mongo
}
