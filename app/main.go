package main

import (
	"Canteen-Backend/internal/delivery"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/internal/repository/postgres"
	"Canteen-Backend/internal/usecase"
	"Canteen-Backend/logger"
	"Canteen-Backend/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	if err := logger.Init(); err != nil {
		log.Fatalf("error initializing logger: %s", err.Error())
	}
}

// @title Canteen Management System API
// @version 1.0
// @host localhost:8080
func main() {

	db, err := postgres.ConnectDatabase(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		log.Fatalf("error occured while connecting to db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	useCase := usecase.NewUseCase(repo)
	handler := delivery.NewHandler(useCase)

	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
