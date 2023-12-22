package config

import (
	"os"

	"github.com/moxicom/todo-back/pkg/repository"
	"github.com/spf13/viper"
)

func InitDbConfig() (repository.Config, error) {
	return repository.Config{
		Host:     viper.GetString("host"),
		User:     viper.GetString("user"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   viper.GetString("dbname"),
		Port:     viper.GetString("port"),
		SSLMode:  viper.GetString("sslmode"),
	}, nil
}
