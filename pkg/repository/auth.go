package repository

import (
	"github.com/moxicom/todo-back/models"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) CreateUser(user models.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func (r *authRepository) CheckUsernameExistence(username string) error {
	var result models.User
	query := "username = ?"
	err := r.db.Where(query, username).First(&result).Error
	return err
}

func (r *authRepository) GetUser(username, password string) (models.User, error) {
	var result models.User
	query := "username = ? AND password = ?"
	err := r.db.Where(query, username, password).First(&result).Error
	return result, err
}
