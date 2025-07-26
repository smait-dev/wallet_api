package main

import (
	"log"
	"wallet_api/internal/handler"
	"wallet_api/internal/server"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	service := 0
	handler := handler.NewHandler(service)
	server := &server.Server{}

	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal("server error")
	}
}

func InitConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}