package main

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wallet_api/internal/config"
	"wallet_api/internal/database"
	"wallet_api/internal/errors"
	"wallet_api/internal/handler"
	"wallet_api/internal/repository"
	"wallet_api/internal/server"
	"wallet_api/internal/service"
)

func main() {
	config := config.InitConfig()

	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal(errors.ErrMigrations)
	}
	err = database.DBSeedWallets(db)
	if err != nil {
		log.Fatal(errors.ErrSeed)
	}

	repository := repository.NewRepository(db, config)
	service := service.NewService(repository, config)
	handler := handler.NewHandler(service, config)
	server := &server.Server{}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// запуск сервера
	go func() {
		if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			log.Fatal(errors.ErrServerRun)
		}
	}()
	// выключение сервера
	<-quit
	log.Println("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("shutdown failed: ", err.Error())
	}
}
