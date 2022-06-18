package handler

import (
	"demo/service"
	"encoding/json"
	"log"
	"net/http"
)

type TaskHandlers struct {
	Service service.TaskService
}

func (th *TaskHandlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	tasks, err := th.Service.GetAllTask()
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Fatal(err)
	}
}
