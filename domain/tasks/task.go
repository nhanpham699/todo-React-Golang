package tasks

import (
	"github.com/nhanpham699/demo/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id,omitempty" validate:"required"`
	Name string             `json:"name" bson:"name" validate:"required"`
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
	Create(dto.CreateTaskRequest) (*dto.HandleResponse, *error)
	Delete(dto.DeleteTaskRequest) (*dto.HandleResponse, *error)
	Update(dto.UpdateTaskRequest) (*dto.HandleResponse, *error)
}
