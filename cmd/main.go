package main

import (
	"log"

	todoback "github.com/moxicom/todo-back"
	"github.com/moxicom/todo-back/pkg/handlers"
)

func main() {
	handlers := handlers.Handler{}
	server := todoback.Server{}

	if err := server.Run("8080", handlers.InitRoutes().Handler()); err != nil {
		log.Fatal(err)
	}
}
