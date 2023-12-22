package service

import (
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

type AuthService struct {
	repository repository.Repository
}

func NewAuthService(r repository.Repository) *AuthService {
	return &AuthService{
		repository: r,
	}
}

func (s *AuthService) CreateUser(user *models.User) (int, error) {
	return s.repository.CreateUser(user)
}
