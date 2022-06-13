package main

import (
	"context"
	"demo/dto"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://nhanpham:nhan123@cluster0.c3e8c.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
	dbName := "todoDB"
	docCollection := "tasks"
	fmt.Print("connnected to MongoDB")
	router := mux.NewRouter()
	address := "localhost"
	port := "8080"
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
	db := client.Database(dbName).Collection(docCollection)
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var tasks []dto.Task
	if err = cur.All(context.TODO(), &tasks); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found multiple documents: %+v\n", tasks)

}
