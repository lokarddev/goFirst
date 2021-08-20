package main

import (
	"github.com/spf13/viper"
	"goFirst"
	"goFirst/pkg/handler"
	"goFirst/pkg/repository"
	"goFirst/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while configuring")
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(goFirst.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
