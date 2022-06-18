package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type TaskReponse struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id,omitempty" validate:"required"`
	Name string             `json:"name" bson:"name" validate:"required"`
}
