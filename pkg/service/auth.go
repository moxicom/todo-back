package service

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

const saltOs = "PASSWORD_SALT"

type AuthService struct {
	repository repository.Auth
}

func NewAuthService(r repository.Auth) *AuthService {
	return &AuthService{
		repository: r,
	}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = makePasswordHash(user.Password)
	return s.repository.CreateUser(user)
}

func makePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := os.Getenv(saltOs)
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
