package handler

import (
	"demo/service"
	"encoding/json"
	"log"
	"net/http"
)

type TaskHandlers struct {
	service service.TaskService
}

func (th *TaskHandlers) getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	tasks, err := th.service.GetAllTask()
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		panic(err)
	}
}
