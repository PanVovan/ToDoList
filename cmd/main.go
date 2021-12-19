package main

import (
	server "ToDoList"
	"ToDoList/pkg/handler"
	"ToDoList/pkg/repository"
	"ToDoList/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializating config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()

}
