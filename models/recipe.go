package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name          string `gorm:"not null;default:null"`
	Description   *string
	Difficulty    *uint
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
	RecipeID        uint
}

type Direction struct {
	gorm.Model
	Description string
	Time        *uint
	Order       uint `gorm:"not null;default:null"`
	RecipeID    uint
}
