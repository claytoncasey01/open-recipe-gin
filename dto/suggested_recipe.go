package dto

import "github.com/claytoncasey01/open-recipe-gin/models"

type SuggestedRecipeDTO struct {
	Name          string                   `json:"name"`
	Description   *string                  `json:"description"`
	Difficulty    *uint                    `json:"difficulty"`
	TotalCalories *uint                    `json:"total_calories"`
	TotalPrepTime *string                  `json:"total_prep_time"`
	Ingredients   []SuggestedIngredientDTO `json:"ingredients"`
	Directions    []SuggestedDirectionDTO  `json:"directions"`
}

type SuggestedIngredientDTO struct {
	Name            string  `json:"name"`
	Quantity        string  `json:"quantity"`
	MeasurementUnit *string `json:"measurement_unit"`
}

type SuggestedDirectionDTO struct {
	Description string `json:"description"`
	Time        *uint  `json:"time"`
	Order       uint   `json:"order"`
}

func SuggestedRecipeDTOFromModel(model models.SuggestedRecipe) SuggestedRecipeDTO {
	suggestedRecipeDTO := SuggestedRecipeDTO{
		Name:          model.Name,
		Description:   model.Description,
		Difficulty:    model.Difficulty,
		TotalCalories: model.TotalCalories,
		TotalPrepTime: model.TotalPrepTime,
	}

	for _, ingredient := range model.Ingredients {
		suggestedRecipeDTO.Ingredients = append(suggestedRecipeDTO.Ingredients, SuggestedIngredientDTO{
			Name:            ingredient.Name,
			Quantity:        ingredient.Quantity,
			MeasurementUnit: ingredient.MeasurementUnit,
		})
	}

	for _, direction := range model.Directions {
		suggestedRecipeDTO.Directions = append(suggestedRecipeDTO.Directions, SuggestedDirectionDTO{
			Description: direction.Description,
			Time:        direction.Time,
			Order:       direction.Order,
		})
	}

	return suggestedRecipeDTO
}

func SuggestedRecipeModelFromDTO(dto SuggestedRecipeDTO) models.SuggestedRecipe {
	suggestedRecipeModel := models.SuggestedRecipe{
		Name:          dto.Name,
		Description:   dto.Description,
		Difficulty:    dto.Difficulty,
		TotalCalories: dto.TotalCalories,
		TotalPrepTime: dto.TotalPrepTime,
	}

	for _, ingredient := range dto.Ingredients {
		suggestedRecipeModel.Ingredients = append(suggestedRecipeModel.Ingredients, models.SuggestedIngredient{
			Name:            ingredient.Name,
			Quantity:        ingredient.Quantity,
			MeasurementUnit: ingredient.MeasurementUnit,
		})
	}

	for _, direction := range dto.Directions {
		suggestedRecipeModel.Directions = append(suggestedRecipeModel.Directions, models.SuggestedDirection{
			Description: direction.Description,
			Time:        direction.Time,
			Order:       direction.Order,
		})
	}

	return suggestedRecipeModel
}

func RecipeModelFromSuggestedRecipeDTO(dto SuggestedRecipeDTO) models.Recipe {
	recipeModel := models.Recipe{
		Name:          dto.Name,
		Description:   dto.Description,
		Difficulty:    dto.Difficulty,
		TotalCalories: dto.TotalCalories,
		TotalPrepTime: dto.TotalPrepTime,
	}

	for _, ingredient := range dto.Ingredients {
		recipeModel.Ingredients = append(recipeModel.Ingredients, models.Ingredient{
			Name:            ingredient.Name,
			Quantity:        ingredient.Quantity,
			MeasurementUnit: ingredient.MeasurementUnit,
		})
	}

	for _, direction := range dto.Directions {
		recipeModel.Directions = append(recipeModel.Directions, models.Direction{
			Description: direction.Description,
			Time:        direction.Time,
			Order:       direction.Order,
		})
	}

	return recipeModel
}
