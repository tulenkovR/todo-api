package main

import (
	"github.com/tulenkovR/todo-api"
	"github.com/tulenkovR/todo-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error %s", err.Error())
	}
}
