package main

import (
	"demo/domain"
	"demo/driver"
	"demo/handler"
	"demo/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	address := "localhost"
	port := "8080"
	taskRepositoryDb := domain.NewTaskRepositoryDb(driver.Mongo.Client)
	th := handler.TaskHandlers{service: service.NewTaskService(taskRepositoryDb)}

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
