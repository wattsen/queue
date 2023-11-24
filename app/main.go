package main

import (
	"banking-system/internal/delivery"
	"banking-system/internal/repository"
	pg "banking-system/internal/repository/postgres"
	"banking-system/internal/server"
	"banking-system/internal/usecase"
	"banking-system/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	if err := logger.Init(); err != nil {
		log.Fatalf("error initializing logger: %s", err.Error())
	}
}

func main() {

	db, err := pg.ConnectDatabase(pg.Config{
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
	go updateCoefficientPeriodically(useCase)
	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func updateCoefficientPeriodically(uc *usecase.UseCase) {
	for {
		time.Sleep(5 * time.Minute)
		err := uc.Ticket.UpdateCoefficient()
		if err != nil {
			log.Printf("error occurred during coefficient update: %s", err.Error())
		}
	}
}
