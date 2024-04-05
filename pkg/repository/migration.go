package repository

import (
	"fmt"

	"github.com/moxicom/todo-back/models"
	"gorm.io/gorm"
)

func NewMigration(db *gorm.DB) error {
	fmt.Println("Making migration...")
	err := db.AutoMigrate(
		&models.Item{},
		&models.TodoList{},
		&models.User{},
		&models.ListItem{},
		&models.UserList{})
	return err
}
