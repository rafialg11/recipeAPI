package usecase

import (
	"github.com/rafialg11/recipe-api/internal/domain"
	"github.com/rafialg11/recipe-api/internal/repository"
)

type RecipeUseCase interface {
	GetAllRecipes() ([]domain.Recipe, error)
	GetRecipeByID(id int) (domain.Recipe, error)
	CreateRecipe(recipe domain.Recipe) (domain.Recipe, error)
}

type recipeUseCase struct {
	recipeRepo repository.RecipeRepository
}

func NewRecipeUseCase(repo repository.RecipeRepository) RecipeUseCase {
	return &recipeUseCase{recipeRepo: repo}
}

func (u *recipeUseCase) GetAllRecipes() ([]domain.Recipe, error) {
	return u.recipeRepo.FetchAll()
}

func (u *recipeUseCase) GetRecipeByID(id int) (domain.Recipe, error) {
	return u.recipeRepo.FetchByID(id)
}

func (u *recipeUseCase) CreateRecipe(recipe domain.Recipe) (domain.Recipe, error) {
	return u.recipeRepo.Save(recipe)
}
