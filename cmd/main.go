package main

import (
	"log"

	"github.com/ilidiance/todo-app"
	"github.com/ilidiance/todo-app/pkg/handler"
	"github.com/ilidiance/todo-app/pkg/repository"
	"github.com/ilidiance/todo-app/pkg/service"
)

func main() {
	//handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running httpserver, %s", err.Error())
	}
}
