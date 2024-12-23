package main

import (
	"database/sql"
	"fmt"
	"log"
	"user/config"
	"user/internal/delivery/http"
	"user/internal/repository"
	"user/internal/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize PostgreSQL connection
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Initialize repository and usecase
	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret)

	// Set up router
	router := gin.Default()

	// Register routes
	http.RegisterAuthRoutes(router, authUsecase)

	// Start server
	serverAddress := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Printf("Starting server on %s", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
