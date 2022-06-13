package domain

import "demo/dto"

type Task struct {
	Id   string `db:"task_id"`
	Name string `db:"task_name"`
}

func (t Task) TaskDto() dto.TaskResponse {
	return dto.TaskResponse{
		Id:   t.Id,
		Name: t.Name,
	}
}

type TaskRepository interface {
	FindAll() ([]Task, *error)
	ById(id string) (*Task, *error)
}
