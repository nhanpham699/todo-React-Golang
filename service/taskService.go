package service

import (
	"demo/dto"
)

type TaskService interface {
	getAllTask(string) ([]dto.Task, *error)
	getTask(string) (*dto.Task, *error)
}

type DefaultTaskService interface {
	repo domain
}
