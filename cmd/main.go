package main

import (
	"log"

	todoback "github.com/moxicom/todo-back"
	"github.com/moxicom/todo-back/pkg/handlers"
)

func main() {
	handlers := new(handlers.Handler)
	server := new(todoback.Server)

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
