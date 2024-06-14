package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name          string
	Description   *string
	Difficulty    *string // Enum probably? TODO: Look back into this
	TotalCalories *uint
	TotalPrepTime *uint
	Ingredients   []Ingredient
	Directions    []Direction
}

type Ingredient struct {
	gorm.Model
	Name            string
	Quantity        string
	MeasurementUnit *string
	Calories        *uint
	RecipeID        uint
}

type Direction struct {
	gorm.Model
	description string
	time        *uint
	order       uint
	RecipeID    uint
}
