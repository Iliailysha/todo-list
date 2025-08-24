package main

import (
	"log"
	todo "todo-list"
	"todo-list/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.Initroutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
