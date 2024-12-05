package usecase

import (
	service "github.com/rafialg11/recipe-api/internal/application/services"
	"github.com/rafialg11/recipe-api/internal/domain"
)

type RecipeUseCase interface {
	GetAllRecipes() ([]domain.Recipe, error)
	GetRecipeByID(id int) (domain.Recipe, error)
	CreateRecipe(recipe domain.Recipe) (domain.Recipe, error)
}

type recipeUseCase struct {
	recipeService service.RecipeService
}

func NewRecipeUseCase(service service.RecipeService) RecipeUseCase {
	return &recipeUseCase{recipeService: service}
}

func (u *recipeUseCase) GetAllRecipes() ([]domain.Recipe, error) {
	return u.recipeService.GetAllRecipes()
}

func (u *recipeUseCase) GetRecipeByID(id int) (domain.Recipe, error) {
	return u.recipeService.GetRecipeByID(id)
}

func (u *recipeUseCase) CreateRecipe(recipe domain.Recipe) (domain.Recipe, error) {
	return u.recipeService.CreateRecipe(recipe)
}
