package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rafialg11/recipe-api/internal/delivery/handler"
	"github.com/rafialg11/recipe-api/internal/usecase"
)

func SetupRoutes(e *echo.Echo, recipeUseCase usecase.RecipeUseCase) {
	recipeHandler := handler.NewRecipeHandler(recipeUseCase)

	e.GET("/recipes", recipeHandler.GetAllRecipes)
	e.POST("/recipes", recipeHandler.CreateRecipe)
}
