package dto

type RecipeDTO struct {
	Title        string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
	Ingredients  string `json:"ingredients"`
}
