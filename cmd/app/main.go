package main

import (
	"Canteen-Backend/cmd/app/server"
	"Canteen-Backend/cmd/migrations"
	"Canteen-Backend/internal/delivery/handlers"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/internal/usecase"
	"Canteen-Backend/internal/utils"
	"Canteen-Backend/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
)

func init() {
	if err := logger.Init(); err != nil {
		log.Fatalf("error initializing logger: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logger.GetLogger().Fatal("error loading env variables", zap.Error(err))
	}
}

// @title Canteen Management System API
// @version 1.0
// @host localhost:8080
// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func main() {

	db, err := utils.ConnectDatabase(utils.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_NAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	})
	if err != nil {
		logger.GetLogger().Fatal("error occurred while connecting to db", zap.Error(err))
	}

	if err := migrations.RunMigrations(db); err != nil {
		logger.GetLogger().Fatal("error occurred while running migrations", zap.Error(err))
	}

	if err := migrations.RunSeeds(db); err != nil {
		logger.GetLogger().Fatal("error occurred while running seeds", zap.Error(err))
	}

	repo := repository.NewRepository(db)
	useCase := usecase.NewUseCase(repo)
	handler := handlers.NewHandler(useCase)

	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logger.GetLogger().Fatal("error occurred while running http server", zap.Error(err))
	}
}
