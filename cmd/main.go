package main

import (
	"log"
	"os"
	todo "todo/study"
	"todo/study/package/handler"
	"todo/study/package/repository"
	"todo/study/package/service"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB(repository.DBConfig{
		Host:     viper.GetString("DB_Config.Host"),
		Port:     viper.GetString("DB_Config.Port"),
		Username: viper.GetString("DB_Config.Username"),
		Password: os.Getenv("DB_PASS"),
		DBName:   viper.GetString("DB_Config.DBName"),
		SSLMode:  viper.GetString("DB_Config.SSLMode"),
	})

	if err != nil {
		log.Fatalf("Can't DB connect: %s", err)
	}

	repos := repository.NewRepository(db)
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
