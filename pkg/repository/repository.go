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
	CheckList(userId, listId int) error
	Create(listId int, item models.Item) (int, error)
	GetAll(listId int) ([]models.Item, error)
	GetById(listId, itemId int) (models.Item, error)
	Delete(listId, itemId int) error
	Update(listId, itemId int, input models.Item) error
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
		Item:     newItemRepository(db),
	}
}
