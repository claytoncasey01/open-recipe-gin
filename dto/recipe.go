package dto

import "github.com/claytoncasey01/open-recipe-gin/models"

type RecipeFilters struct {
	Name       string `form:"name"`
	Difficulty string `form:"difficulty"`
	PrepTime   uint   `form:"prepTime"`
}

type RecipeDTO struct {
	ID            uint            `json:"id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description"`
	Difficulty    *uint           `json:"difficulty"`
	TotalCalories *uint           `json:"total_calories"`
	TotalPrepTime *string         `json:"total_prep_time"`
	Ingredients   []IngredientDTO `json:"ingredients"`
	Directions    []DirectionDTO  `json:"directions"`
}

type IngredientDTO struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Quantity        string  `json:"quantity"`
	MeasurementUnit *string `json:"measurement_unit"`
}

type DirectionDTO struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Time        *uint  `json:"time"`
	Order       uint   `json:"order"`
}

func RecipeDTOFromModel(model models.Recipe) RecipeDTO {
	recipeDTO := RecipeDTO{
		ID:            model.ID,
		Name:          model.Name,
		Description:   model.Description,
		Difficulty:    model.Difficulty,
		TotalCalories: model.TotalCalories,
		TotalPrepTime: model.TotalPrepTime,
	}

	for _, ingredient := range model.Ingredients {
		recipeDTO.Ingredients = append(recipeDTO.Ingredients, IngredientDTO{
			ID:              ingredient.ID,
			Name:            ingredient.Name,
			Quantity:        ingredient.Quantity,
			MeasurementUnit: ingredient.MeasurementUnit,
		})
	}

	for _, direction := range model.Directions {
		recipeDTO.Directions = append(recipeDTO.Directions, DirectionDTO{
			ID:          direction.ID,
			Description: direction.Description,
			Time:        direction.Time,
			Order:       direction.Order,
		})
	}

	return recipeDTO
}

func RecipeModelFromDTO(dto RecipeDTO) models.Recipe {
	recipeModel := models.Recipe{
		ID:            dto.ID,
		Name:          dto.Name,
		Description:   dto.Description,
		Difficulty:    dto.Difficulty,
		TotalCalories: dto.TotalCalories,
		TotalPrepTime: dto.TotalPrepTime,
	}

	for _, ingredient := range dto.Ingredients {
		recipeModel.Ingredients = append(recipeModel.Ingredients, models.Ingredient{
			ID:              ingredient.ID,
			Name:            ingredient.Name,
			Quantity:        ingredient.Quantity,
			MeasurementUnit: ingredient.MeasurementUnit,
		})
	}

	for _, direction := range dto.Directions {
		recipeModel.Directions = append(recipeModel.Directions, models.Direction{
			ID:          direction.ID,
			Description: direction.Description,
			Time:        direction.Time,
			Order:       direction.Order,
		})
	}

	return recipeModel
}
