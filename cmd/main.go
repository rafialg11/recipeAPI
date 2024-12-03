package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rafialg11/recipe-api/internal/config"

	"github.com/rafialg11/recipe-api/internal/delivery/router"
	"github.com/rafialg11/recipe-api/internal/repository"
	"github.com/rafialg11/recipe-api/internal/usecase"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the Echo instance
	e := echo.New()

	// Initialize the repository with DB connection
	recipeRepo, err := repository.NewRecipeRepository(cfg.DbURL)
	if err != nil {
		log.Fatalf("Could not initialize repository: %v", err)
	}

	// Initialize use case
	recipeUseCase := usecase.NewRecipeUseCase(recipeRepo)

	// Set up routes
	router.SetupRoutes(e, recipeUseCase)

	// Start the server
	port := ":" + cfg.ServerPort
	log.Printf("Starting server on port %s...\n", port)
	if err := e.Start(port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
