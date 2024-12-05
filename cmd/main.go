package main

import (
	"log"

	"github.com/labstack/echo/v4"
	service "github.com/rafialg11/recipe-api/internal/application/services"
	"github.com/rafialg11/recipe-api/internal/config"
	"github.com/rafialg11/recipe-api/internal/database"
	"github.com/rafialg11/recipe-api/internal/delivery/router"
	"github.com/rafialg11/recipe-api/internal/infrastructure/repository"
	"github.com/rafialg11/recipe-api/internal/usecase"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize the database connection
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	// Initialize the repository
	recipeRepo := repository.NewRecipeRepository(db)

	// Initialize the service
	recipeService := service.NewRecipeService(recipeRepo)

	// Initialize the use case
	recipeUseCase := usecase.NewRecipeUseCase(recipeService)

	// Initialize the Echo instance
	e := echo.New()

	// Set up routes with the use case and handler
	router.SetupRoutes(e, recipeUseCase)

	// Start the server
	port := ":" + cfg.ServerPort
	log.Printf("Starting server on port %s...\n", port)
	if err := e.Start(port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
