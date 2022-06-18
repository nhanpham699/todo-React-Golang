package domain

import (
	"demo/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id,omitempty" validate:"required"`
	Name string             `json:"name" bson:"name" validate:"required"`
}

func (t Task) TaskDto() dto.TaskReponse {
	return dto.TaskReponse{
		Id:   t.Id,
		Name: t.Name,
	}
}

type TaskRepository interface {
	FindAll() ([]Task, *error)
	// ById(id string) (*Task, *error)
}
