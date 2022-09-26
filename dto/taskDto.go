package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type TaskResponse struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id,omitempty" validate:"required"`
	Name string             `json:"name" bson:"name" validate:"required"`
}

type CreateTaskRequest struct {
	Name string `json:"name" bson:"name" validate:"required"`
}

type UpdateTaskRequest struct {
	Id   string `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type HandleResponse struct {
	ResultCode string `json:"resultCode" bson:"resultCode`
	Message    string `json:"message" bson:"message"`
}

type DeleteTaskRequest struct {
	Id string `json:"_id" bson:"_id"`
}
