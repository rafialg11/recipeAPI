package main

import (
	"log"

	"github.com/labstack/echo/v4"
	// "github.com/rafialg11/recipe-api/internal/delivery/router"
	// "github.com/rafialg11/recipe-api/internal/config"
)

func main() {
	// cfg :=

	//Init Echo instance
	e := echo.New()

	//Start the server
	port := ":" + cfg.ServerPort
	log.Printf("Starting server on port %s", port)
	if err := e.Start(port); err != nil {
		log.Fatal("Error starting server: ", err)

	}

}
