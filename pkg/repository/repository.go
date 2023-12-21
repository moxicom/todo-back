package repository

import "gorm.io/gorm"

// DB handlers

type Auth interface {
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
	return &Repository{}
}
