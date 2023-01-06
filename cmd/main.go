package main

import (
	"log"
	todo "todo/study"
	"todo/study/package/handler"
	"todo/study/package/repository"
	"todo/study/package/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)

	err := srv.Run("2002", handler.InitRoutes())
	if err != nil {
		log.Fatalf("error %s", err.Error())
	}
}
