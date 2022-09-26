package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nhanpham699/demo/domain/tasks"
	"github.com/nhanpham699/demo/domain/user"
	"github.com/nhanpham699/demo/driver"
	"github.com/nhanpham699/demo/handler"
	"github.com/nhanpham699/demo/service"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	address := "localhost"
	port := "8080"
	taskRepositoryDb := tasks.NewTaskRepositoryDb(driver.ConnectMongoDB().Client)
	th := handler.TaskHandlers{service.NewTaskService(taskRepositoryDb)}

	userRepositoryDb := user.NewUserRepositoryDb(driver.ConnectMongoDB().Client)
	uh := handler.UserHandlers{service.NewUserService(userRepositoryDb)}

	router.
		HandleFunc("/tasks", th.GetAllTasks).
		Methods(http.MethodGet).
		Name("GetAllTasks")
	router.
		HandleFunc("/tasks/{task_id}", th.GetTask).
		Methods(http.MethodGet).
		Name("GetTaskById")
	router.
		HandleFunc("/tasks/create", th.CreateTask).
		Methods(http.MethodPost).
		Name("CreateTask")
	router.
		HandleFunc("/tasks/update", th.UpdateTask).
		Methods(http.MethodPost).
		Name("UpdateTask")
	router.
		HandleFunc("/tasks/delete", th.DeleteTask).
		Methods(http.MethodPost).
		Name("DeleteTask")

	router.
		HandleFunc("/users/register", uh.Register).
		Methods(http.MethodPost).
		Name("RegisterUser")
	router.
		HandleFunc("/users/login", uh.Login).
		Methods(http.MethodPost).
		Name("LoginUser")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
