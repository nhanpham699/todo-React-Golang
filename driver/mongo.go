package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var user = "nhanpham"
var password = "nhan123"
var mongoUri = "mongodb+srv://" + user + ":" + password + "@cluster0.c3e8c.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

var Mongo = &MongoDB{}

func mongoDB() *MongoDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Print("connnected to MongoDB!")
	Mongo.Client = client
	return Mongo
}
