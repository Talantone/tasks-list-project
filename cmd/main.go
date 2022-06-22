package main

import (
	"log"
	tasks_app "tasks-list-project"
	"tasks-list-project/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	server := new(tasks_app.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}
