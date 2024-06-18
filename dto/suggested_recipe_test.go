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
		ID:            1,
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []models.SuggestedIngredient{
			{ID: 1, Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []models.SuggestedDirection{
			{ID: 1, Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	expected := SuggestedRecipeDTO{
		ID:            1,
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []SuggestedIngredientDTO{
			{ID: 1, Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []SuggestedDirectionDTO{
			{ID: 1, Description: "Mix ingredients", Time: nil, Order: 1},
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
		ID:            2,
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []SuggestedIngredientDTO{
			{ID: 2, Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []SuggestedDirectionDTO{
			{ID: 2, Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	expected := models.SuggestedRecipe{
		ID:            2,
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []models.SuggestedIngredient{
			{ID: 2, Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []models.SuggestedDirection{
			{ID: 2, Description: "Mix ingredients", Time: nil, Order: 1},
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
		ID:            3,
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []SuggestedIngredientDTO{
			{ID: 3, Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []SuggestedDirectionDTO{
			{ID: 3, Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	expected := models.Recipe{
		ID:            3,
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []models.Ingredient{
			{ID: 3, Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []models.Direction{
			{ID: 3, Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	result := RecipeModelFromSuggestedRecipeDTO(dto)

	assert.Equal(t, expected, result, "RecipeModelFromSuggestedRecipeDTO() should convert DTO to model correctly")
}
