package main

import (
	"context"
	"flag"

	"github.com/vet-clinic-back/info-service/internal/config"
	"github.com/vet-clinic-back/info-service/internal/handlers"
	"github.com/vet-clinic-back/info-service/internal/logging"
	"github.com/vet-clinic-back/info-service/internal/server"
	"github.com/vet-clinic-back/info-service/internal/service"
	"github.com/vet-clinic-back/info-service/internal/storage"
)

//  @title      Vet clinic auth service
//  @version    0.1
//  @description  auth service

//  @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in              header
// @name            Authorization
func main() {
	isLocal := flag.Bool("local", false, "is it local? can make logs pretty")
	idDebug := flag.Bool("debug", false, "is it local? can make logs pretty")
	flag.Parse()

	log := logging.NewLogger(isLocal, idDebug)
	log.Info("logger initialized")

	log.Info("initializing config")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Failed to load config. ", err)
	}

	log.Info("initializing storage")
	storage := storage.New(log, &cfg.Db)
	defer storage.StorageProcess.Shutdown()

	log.Info("initializing service")
	service := service.New(log, storage.Info)

	log.Info("initializing handler")
	hander := handlers.NewHandler(log, service)

	log.Info("initializing server")
	server := server.NewServer()

	log.Info("starting server on port 8080")
	server.Run("8080", hander.InitRoutes()) // TODO use port from config

	log.Info("graceful shutdown")
	server.Shutdown(context.Background())
}
