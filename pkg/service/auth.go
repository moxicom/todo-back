package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

const (
	saltOs     = "PASSWORD_SALT"
	signingKey = "wesdrftgyhujikujgh"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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
	err := s.repository.CheckUsernameExistence(user.Username)
	if err == nil {
		return 0, errors.New(fmt.Sprintf("User with username %s is already exists", user.Username))
	}
	return s.repository.CreateUser(user)
}

func makePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := os.Getenv(saltOs)
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CreateToken(username, password string) (string, error) {
	user, err := s.repository.GetUser(username, makePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
