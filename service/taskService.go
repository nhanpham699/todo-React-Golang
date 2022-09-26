package service

import (
	domain "github.com/nhanpham699/demo/domain/tasks"
	"github.com/nhanpham699/demo/dto"
)

type TaskService interface {
	GetAll() ([]dto.TaskResponse, *error)
	GetById(string) (*dto.TaskResponse, *error)
	Create(dt dto.CreateTaskRequest) (*dto.HandleResponse, *error)
	Delete(dt dto.DeleteTaskRequest) (*dto.HandleResponse, *error)
	Update(dt dto.UpdateTaskRequest) (*dto.HandleResponse, *error)
}

type DefaultTaskService struct {
	repo domain.TaskRepository
}

func (s DefaultTaskService) GetAll() ([]dto.TaskResponse, *error) {
	tasks, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.TaskResponse, 0)
	for _, c := range tasks {
		response = append(response, c.TaskDto())
	}
	return response, err
}

func (s DefaultTaskService) GetById(id string) (*dto.TaskResponse, *error) {
	tasks, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := tasks.TaskDto()
	return &response, err
}

func (s DefaultTaskService) Create(data dto.CreateTaskRequest) (*dto.HandleResponse, *error) {
	response, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (s DefaultTaskService) Delete(data dto.DeleteTaskRequest) (*dto.HandleResponse, *error) {
	response, err := s.repo.Delete(data)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (s DefaultTaskService) Update(data dto.UpdateTaskRequest) (*dto.HandleResponse, *error) {
	response, err := s.repo.Update(data)
	if err != nil {
		return nil, err
	}

	return response, err
}

func NewTaskService(repository domain.TaskRepository) DefaultTaskService {
	return DefaultTaskService{repository}
}
