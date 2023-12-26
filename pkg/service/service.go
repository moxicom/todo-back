package service

import (
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

// Business logic

type Auth interface {
	CreateUser(user models.User) (int, error)
	CreateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
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
		Auth:     NewAuthService(repository.Auth),
		TodoList: NewTodoService(repository.TodoList),
	}
}
