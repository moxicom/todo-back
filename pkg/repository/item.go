package repository

import (
	"github.com/moxicom/todo-back/models"
	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func newItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{
		db: db,
	}
}

func (r *itemRepository) CheckList(userId, listId int) error {
	return r.db.Where("user_id = ? AND list_id = ?", userId, listId).First(&models.UserList{}).Error
}

func (r *itemRepository) Create(listId int, item models.Item) (int, error) {
	tx := r.db.Begin()

	if err := tx.Create(&item).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	itemList := models.ListItem{
		ListId: listId,
		ItemId: item.Id,
	}
	if err := tx.Create(&itemList).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return item.Id, tx.Commit().Error
}

func (r *itemRepository) GetAll(listId int) ([]models.Item, error) {
	var items []models.Item

	err := r.db.
		Joins("INNER JOIN list_items li ON items.id = li.item_id").
		Where("li.list_id = ?", listId).Find(&items).Error
	if err != nil {
		return []models.Item{}, err
	}

	return items, nil
}

func (r *itemRepository) GetById(listId, itemId int) (models.Item, error) {
	var item models.Item
	err := r.db.
		Joins("INNER JOIN list_items li ON items.id = li.item_id").
		Where("li.list_id = ? AND items.id = ?", listId, itemId).First(&item).Error
	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (r *itemRepository) Delete(itemId int) error {
	return nil
}
func (r *itemRepository) Update(listId, itemId int, input models.Item) error {
	item, err := r.GetById(listId, itemId)
	if err != nil {
		return err
	}

	item.Title = input.Title
	item.Description = input.Description
	item.Done = input.Done

	err = r.db.Save(&item).Error
	if err != nil {
		return err
	}

	return nil
}
