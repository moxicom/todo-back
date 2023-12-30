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

func (r *todoListRepository) Update(userId, listId int, input models.TodoList) error {
	list, err := r.GetById(userId, listId)
	if err != nil {
		return err
	}

	list.Title = input.Title
	list.Description = input.Description

	err = r.db.Model(&list).Updates(models.TodoList{Title: input.Title, Description: input.Description}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *todoListRepository) Delete(userId, listId int) error {
	_, err := r.GetById(userId, listId)
	if err != nil {
		return err
	}

	// delete
	tx := r.db.Begin()
	if err := tx.Where("user_id = ? AND list_id = ?", userId, listId).Delete(&models.UserList{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", listId).Delete(&models.TodoList{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
