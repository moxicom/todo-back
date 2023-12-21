package main

import (
	"fmt"
	"log"

	todoback "github.com/moxicom/todo-back"
	config "github.com/moxicom/todo-back/configs"
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/handlers"
	"github.com/moxicom/todo-back/pkg/repository"
	"github.com/moxicom/todo-back/pkg/service"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Config init
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Db opening
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("host"),
		viper.GetString("user"),
		viper.GetString("password"),
		viper.GetString("dbname"),
		viper.GetString("port"),
		viper.GetString("sslmode"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	//Db Migration
	err = db.AutoMigrate(&models.ListItem{}, &models.UserList{}, &models.Item{}, &models.User{}, &models.TodoList{})
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Dependency injection
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)
	server := todoback.NewServer()

	if err = server.Run(viper.GetString("server_port"), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
