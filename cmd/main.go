package main

import (
	"log"

	todoback "github.com/moxicom/todo-back"
	"github.com/moxicom/todo-back/pkg/handlers"
	"github.com/moxicom/todo-back/pkg/repository"
	"github.com/moxicom/todo-back/pkg/service"
)

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)
	server := todoback.NewServer()

	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
