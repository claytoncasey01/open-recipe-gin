package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null;default:null"`
	Description   *string
	Difficulty    *uint `gorm:"check:difficulty >= 1 AND difficulty <= 10"`
	TotalCalories *uint
	TotalPrepTime *string
	Ingredients   []Ingredient
	Directions    []Direction
}

type Ingredient struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"not null;default:null"`
	Quantity        string
	MeasurementUnit *string
	RecipeID        uint
}

type Direction struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Description string
	Time        *uint
	Order       uint `gorm:"not null;default:null"`
	RecipeID    uint
}
