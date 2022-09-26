package service

import (
	domain "github.com/nhanpham699/demo/domain/user"
	"github.com/nhanpham699/demo/dto"
)

type UserService interface {
	Login(user dto.AuthRequest) (*dto.AuthResponse, *error)
	Register(user dto.RegisterRequest) (*dto.HandleResponse, *error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) Register(data dto.RegisterRequest) (*dto.HandleResponse, *error) {
	response, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (s DefaultUserService) Login(data dto.AuthRequest) (*dto.AuthResponse, *error) {
	response, err := s.repo.Login(data)
	if err != nil {
		return nil, err
	}
	return response, err
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
