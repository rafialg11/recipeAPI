package domain

type Recipe struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Instructions string   `json:"instructions"`
	Ingredients  []string `json:"ingredients"`
}
