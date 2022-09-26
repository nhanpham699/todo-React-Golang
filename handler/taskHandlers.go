package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nhanpham699/demo/dto"
	"github.com/nhanpham699/demo/service"

	"github.com/gorilla/mux"
)

type TaskHandlers struct {
	Service service.TaskService
}

func (th *TaskHandlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	tasks, err := th.Service.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Fatal(err)
	}
}

func (th *TaskHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	tasks, err := th.Service.GetById(vars["task_id"])
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Fatal(err)
	}
}

func (th *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var request dto.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Println(request)
	res, error := th.Service.Create(request)
	if error != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
	}
}

func (th *TaskHandlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var request dto.UpdateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	res, error := th.Service.Update(request)
	if error != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
	}
}

func (th *TaskHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var request dto.DeleteTaskRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Println(request)
	res, error := th.Service.Delete(request)
	if error != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
	}
}
