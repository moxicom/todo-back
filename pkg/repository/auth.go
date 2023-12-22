package repository

import (
	"github.com/moxicom/todo-back/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
func (r *AuthRepository) CreateUser(user models.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}
