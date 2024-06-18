package dto

import (
	"testing"

	"github.com/claytoncasey01/open-recipe-gin/models"
	"github.com/stretchr/testify/assert"
)

func TestRecipeDTOFromModel(t *testing.T) {
	description := "Delicious recipe"
	difficulty := uint(2)
	totalCalories := uint(500)
	totalPrepTime := "30m"
	model := models.Recipe{
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

	expected := RecipeDTO{
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []IngredientDTO{
			{Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []DirectionDTO{
			{Description: "Mix ingredients", Time: nil, Order: 1},
		},
	}

	result := RecipeDTOFromModel(model)

	assert.Equal(t, expected, result, "RecipeDTOFromModel() should convert model to DTO correctly")
}

func TestRecipeModelFromDTO(t *testing.T) {
	description := "Delicious recipe"
	difficulty := uint(2)
	totalCalories := uint(500)
	totalPrepTime := "30m"
	dto := RecipeDTO{
		Name:          "Test Recipe",
		Description:   &description,
		Difficulty:    &difficulty,
		TotalCalories: &totalCalories,
		TotalPrepTime: &totalPrepTime,
		Ingredients: []IngredientDTO{
			{Name: "Salt", Quantity: "1 tsp", MeasurementUnit: nil},
		},
		Directions: []DirectionDTO{
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

	result := RecipeModelFromDTO(dto)

	assert.Equal(t, expected, result, "RecipeModelFromDTO() should convert DTO to model correctly")
}
