package dto

type RecipeFilters struct {
	Name       string `form:"name"`
	Difficulty string `form:"difficulty"`
	PrepTime   uint   `form:"prepTime"`
}
