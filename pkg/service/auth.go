package service

import (
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

type AuthService struct {
	repository repository.Auth
}

func NewAuthService(r repository.Auth) *AuthService {
	return &AuthService{
		repository: r,
	}
}

func (s *AuthService) CreateUser(user *models.User) (int, error) {
	return s.repository.CreateUser(user)
}

func (s *AuthService) makePasswordHash(password string) (string, error) {
	return "", nil
}
