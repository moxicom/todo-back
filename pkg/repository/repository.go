package repository

import (
	"github.com/moxicom/todo-back/models"
	"gorm.io/gorm"
)

// DB handlers

type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUser(user models.User) (models.User, error)
}

type TodoList interface {
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
		Auth: NewAuthRepository(db),
	}
}
