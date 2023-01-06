package main

import (
	"log"
	todo "todo/study"
	"todo/study/package/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)

	err := srv.Run("2002", handler.InitRoutes())
	if err != nil {
		log.Fatalf("error %s", err.Error())
	}
}
