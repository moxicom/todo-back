package service

import (
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

// Business logic

type Auth interface {
	CreateUser(user *models.User) (int, error)
}

type TodoList interface {
}

type Item interface {
}

type Service struct {
	Auth
	TodoList
	Item
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repository.Auth),
	}
}
