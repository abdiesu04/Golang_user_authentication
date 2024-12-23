package main

import (
	"log"
	"user/config"
	"user/internal/delivery/http"
	"user/internal/repository"
	"user/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize repository
	db, err := repository.ConnectPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret)

	// Initialize HTTP server
	r := gin.Default()
	http.RegisterAuthRoutes(r, authUsecase)

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
