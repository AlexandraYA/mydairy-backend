package service

import (
	dairy "github.com/alexandraya/mydairy-backend"
	"github.com/alexandraya/mydairy-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user dairy.User) (int, error)
}

type TasksList interface {
}

type Service struct {
	Authorization
	TasksList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
