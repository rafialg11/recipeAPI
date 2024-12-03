package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rafialg11/recipe-api/internal/domain"
	"github.com/rafialg11/recipe-api/internal/usecase"
)

type RecipeHandler struct {
	UseCase usecase.RecipeUseCase
}

func NewRecipeHandler(useCase usecase.RecipeUseCase) *RecipeHandler {
	return &RecipeHandler{UseCase: useCase}
}

func (h *RecipeHandler) GetAllRecipes(c echo.Context) error {
	recipes, err := h.UseCase.GetAllRecipes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching recipes")
	}
	return c.JSON(http.StatusOK, recipes)
}

func (h *RecipeHandler) CreateRecipe(c echo.Context) error {
	var recipe domain.Recipe
	if err := c.Bind(&recipe); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	createdRecipe, err := h.UseCase.CreateRecipe(recipe)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating recipe")
	}

	return c.JSON(http.StatusCreated, createdRecipe)
}
