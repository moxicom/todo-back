package repository

import (
	"github.com/moxicom/todo-back/models"
	"gorm.io/gorm"
)

func NewMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Item{},
		&models.TodoList{},
		&models.User{},
		&models.ListItem{},
		&models.UserList{})
	return err
}
