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
	Update(userId, listId int, input models.TodoList) error
	Delete(userId, listId int) error
}

type Item interface {
	CheckList(userId, listId int) error
	Create(userId, listId int, item models.Item) (int, error)
	GetAll(userId, listId int) ([]models.Item, error)
	GetById(userId, listId, itemId int) (models.Item, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.Item) error
}

type Service struct {
	Auth
	TodoList
	Item
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth:     newAuthService(repository.Auth),
		TodoList: newTodoService(repository.TodoList),
		Item:     newItemService(repository.Item),
	}
}
