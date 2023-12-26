package repository

import (
	"fmt"

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
	fmt.Println(userId)

	err := r.db.
		Joins("INNER JOIN user_lists ul ON todo_lists.id = ul.list_id").
		Where("ul.user_id = ?", userId).Find(&lists).
		Error
	if err != nil {
		return []models.TodoList{}, err
	}

	return lists, nil
}

func (r *todoListRepository) GetById(userId, listId int) (models.TodoList, error) {
	var list models.TodoList
	query := r.db.Joins("INNER JOIN user_lists ul ON todo_lists.id = ul.list_id").
		Where("ul.user_id = ? AND ul.list_id = ?", userId, listId).
		First(&list)

	if query.Error != nil {
		return models.TodoList{}, query.Error
	}

	return list, nil
}
