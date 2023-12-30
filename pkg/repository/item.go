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

	// r.db.Joins("INNER JOIN user_lists ul ON todo_lists.id = ul.list_id").
	// 	Where("ul.user_id = ? AND ul.list_id = ?", userId, listId).
	// 	First(&list)
	err := r.db.
		Joins("INNER JOIN list_items li ON items.id = li.item_id").
		Where("li.list_id = ?", listId).Find(&items).Error
	if err != nil {
		return []models.Item{}, err
	}

	return items, nil
}

func (r *itemRepository) GetById(itemId int) (models.Item, error) {
	return models.Item{}, nil
}
func (r *itemRepository) Delete(itemId int) error {
	return nil
}
func (r *itemRepository) Update(itemId int, input models.Item) error {
	return nil
}
