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

func (r *AuthRepository) CheckUsernameExistence(username string) error {
	var result models.User
	query := "username = ?"
	err := r.db.Where(query, username).First(&result).Error
	return err
}

func (r *AuthRepository) GetUser(username, password string) (models.User, error) {
	var result models.User
	query := "username = ? AND password = ?"
	err := r.db.Where(query, username, password).First(&result).Error
	return result, err
}
