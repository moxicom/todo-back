package repository

import (
	"github.com/moxicom/todo-back/models"
	"gorm.io/gorm"
)

// DB handlers

type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
	CheckUsernameExistence(username string) error
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Update(userId, listId int, input models.TodoList) error
	Delete(userId, listId int) error
}

type Item interface {
}

type Repository struct {
	Auth
	TodoList
	Item
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth:     NewAuthRepository(db),
		TodoList: newTodoListRepository(db),
	}
}
