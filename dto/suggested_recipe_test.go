package dto

import (
	"testing"

	"github.com/claytoncasey01/open-recipe-gin/models"
	"github.com/stretchr/testify/assert"
)

func TestSuggestedRecipeDTOFromModel(t *testing.T) {
	description := "Delicious recipe"
	difficulty := uint(2)
	totalCalories := uint(500)
	totalPrepTime := "30m"
	model := models.SuggestedRecipe{
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []models.SuggestedIngredient{
			{Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []models.SuggestedDirection{
			{Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	expected := SuggestedRecipeDTO{
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []SuggestedIngredientDTO{
			{Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []SuggestedDirectionDTO{
			{Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	result := SuggestedRecipeDTOFromModel(model)

	assert.Equal(t, expected, result, "SuggestedRecipeDTOFromModel() should convert model to DTO correctly")
}

func TestSuggestedRecipeModelFromDTO(t *testing.T) {
	description := "Delicious recipe"
	difficulty := uint(2)
	totalCalories := uint(500)
	totalPrepTime := "30m"
	dto := SuggestedRecipeDTO{
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []SuggestedIngredientDTO{
			{Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []SuggestedDirectionDTO{
			{Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	expected := models.SuggestedRecipe{
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []models.SuggestedIngredient{
			{Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []models.SuggestedDirection{
			{Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	result := SuggestedRecipeModelFromDTO(dto)

	assert.Equal(t, expected, result, "SuggestedRecipeModelFromDTO() should convert DTO to model correctly")
}

func TestRecipeModelFromSuggestedRecipeDTO(t *testing.T) {
	description := "Delicious recipe"
	difficulty := uint(2)
	totalCalories := uint(500)
	totalPrepTime := "30m"
	dto := SuggestedRecipeDTO{
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []SuggestedIngredientDTO{
			{Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []SuggestedDirectionDTO{
			{Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	expected := models.Recipe{
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []models.Ingredient{
			{Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []models.Direction{
			{Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	result := RecipeModelFromSuggestedRecipeDTO(dto)

	assert.Equal(t, expected, result, "RecipeModelFromSuggestedRecipeDTO() should convert DTO to model correctly")
}
