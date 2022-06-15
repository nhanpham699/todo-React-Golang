package domain

import "demo/dto"

type Task struct {
	Id   string `db:"_id"`
	Name string `db:"_name"`
}

func (t Task) TaskDto() dto.Task {
	return dto.Task{
		Id:   t.Id,
		Name: t.Name,
	}
}

type TaskRepository interface {
	FindAll() ([]Task, *error)
	// ById(id string) (*Task, *error)
}
