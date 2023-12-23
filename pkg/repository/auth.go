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

func (r *AuthRepository) GetUser(user models.User) (models.User, error) {
	var result models.User
	query := "username = ? AND password = ?"
	queryRes := r.db.Where(query, user.Username, user.Password)
	return result, queryRes.Error
}
