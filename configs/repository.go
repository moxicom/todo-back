package config

import (
	"os"

	"github.com/moxicom/todo-back/pkg/repository"
)

func InitDbConfig() (repository.Config, error) {
	return repository.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Port:     "5432",
		SSLMode:  os.Getenv("SSL_MODE"),
	}, nil
}
