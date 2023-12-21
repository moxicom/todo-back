package main

import (
	"log"

	todoback "github.com/moxicom/todo-back"
	config "github.com/moxicom/todo-back/configs"
	"github.com/moxicom/todo-back/pkg/handlers"
	"github.com/moxicom/todo-back/pkg/repository"
	"github.com/moxicom/todo-back/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	// Config init
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Db opening
	cfg, err := config.InitDbConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	db, err := repository.NewDbInit(cfg)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	//Db Migration
	err = repository.NewMigration(db)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Dependency injection
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)
	server := todoback.NewServer()

	if err = server.Run(viper.GetString("server_port"), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
