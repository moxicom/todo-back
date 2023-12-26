package repository

import (
	"github.com/moxicom/todo-back/models"
	"gorm.io/gorm"
)

type todoListRepository struct {
	db *gorm.DB
}

func newTodoListRepository(db *gorm.DB) *todoListRepository {
	return &todoListRepository{
		db: db,
	}
}

func (r *todoListRepository) Create(userId int, list models.TodoList) (int, error) {
	tx := r.db.Begin()

	listTemp := models.TodoList{
		Title:       list.Title,
		Description: list.Description,
	}

	if err := tx.Create(&listTemp).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	userList := models.UserList{
		UserId: userId,
		ListId: listTemp.Id,
	}
	if err := tx.Create(&userList).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return listTemp.Id, tx.Commit().Error
}

func (r *todoListRepository) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	query := "id = ?"
	err := r.db.Where(query, userId).Find(&lists).Error
	if err != nil {
		return []models.TodoList{}, err
	}

	return lists, nil
}
