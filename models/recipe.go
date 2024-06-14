package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name          string `gorm:"not null;default:null"`
	Description   *string
	Difficulty    *string // Enum probably? TODO: Look back into this
	TotalCalories *uint
	TotalPrepTime *uint
	Ingredients   []Ingredient
	Directions    []Direction
}

type Ingredient struct {
	gorm.Model
	Name            string `gorm:"not null;default:null"`
	Quantity        string
	MeasurementUnit *string
	Calories        *uint
	RecipeID        uint
}

type Direction struct {
	gorm.Model
	Description string
	Time        *uint
	Order       uint `gorm:"not null;default:null"`
	RecipeID    uint
}
