package service

import (
	"github.com/rafialg11/recipe-api/internal/domain"
	"github.com/rafialg11/recipe-api/internal/repository"
)

type RecipeService interface {
	GetAllRecipes() ([]domain.Recipe, error)
	GetRecipeByID(id int) (domain.Recipe, error)
	CreateRecipe(recipe domain.Recipe) (domain.Recipe, error)
}

type recipeService struct {
	repo repository.RecipeRepository
}

func NewRecipeService(repo repository.RecipeRepository) RecipeService {
	return &recipeService{repo: repo}
}

func (s *recipeService) GetAllRecipes() ([]domain.Recipe, error) {
	// Add any processing logic or filtering here if needed
	return s.repo.FetchAll()
}

func (s *recipeService) GetRecipeByID(id int) (domain.Recipe, error) {
	// Add any processing logic or validation here if needed
	return s.repo.FetchByID(id)
}

func (s *recipeService) CreateRecipe(recipe domain.Recipe) (domain.Recipe, error) {
	// Add any processing or validation logic here if needed
	return s.repo.Save(recipe)
}
