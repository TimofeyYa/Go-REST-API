package main

import (
	"log"
	todo "todo/study"
	"todo/study/package/handler"
	"todo/study/package/repository"
	"todo/study/package/service"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error config: %s", err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("error start server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	return viper.ReadInConfig()
}
