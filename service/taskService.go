package service

import (
	"demo/domain"
	"demo/dto"
)

type TaskService interface {
	GetAllTask() ([]dto.Task, *error)
	// GetTask(string) (*dto.Task, *error)
}

type DefaultTaskService struct {
	repo domain.TaskRepository
}

func (s DefaultTaskService) GetAllTask() ([]dto.Task, *error) {
	tasks, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.Task, 0)
	for _, c := range tasks {
		response = append(response, c.TaskDto())
	}
	return response, err
}

func NewTaskService(repository domain.TaskRepository) DefaultTaskService {
	return DefaultTaskService{repository}
}
