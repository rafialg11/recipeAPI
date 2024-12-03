package repository

import (
	"github.com/jinzhu/gorm" // Or github.com/gorm.io/gorm for GORM v2
	"github.com/rafialg11/recipe-api/internal/domain"
)

type RecipeRepository interface {
	FetchAll() ([]domain.Recipe, error)
	FetchByID(id int) (domain.Recipe, error)
	Save(recipe domain.Recipe) (domain.Recipe, error)
}

type recipeRepository struct {
	DB *gorm.DB
}

// NewRecipeRepository initializes the repository with an existing database connection
func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{DB: db}
}

// FetchAll retrieves all recipes from the database
func (r *recipeRepository) FetchAll() ([]domain.Recipe, error) {
	var recipes []domain.Recipe
	if err := r.DB.Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

// FetchByID retrieves a recipe by its ID
func (r *recipeRepository) FetchByID(id int) (domain.Recipe, error) {
	var recipe domain.Recipe
	if err := r.DB.First(&recipe, id).Error; err != nil {
		return domain.Recipe{}, err
	}
	return recipe, nil
}

// Save persists a new recipe to the database
func (r *recipeRepository) Save(recipe domain.Recipe) (domain.Recipe, error) {
	if err := r.DB.Create(&recipe).Error; err != nil {
		return domain.Recipe{}, err
	}
	return recipe, nil
}
