package user

import "github.com/nhanpham699/demo/dto"

type User struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type UserRepository interface {
	Login(dto.AuthRequest) (*dto.AuthResponse, *error)
	Create(dto.RegisterRequest) (*dto.HandleResponse, *error)
}
