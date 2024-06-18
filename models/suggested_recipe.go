package models

import "gorm.io/gorm"

type SuggestedRecipe struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null;default:null"`
	Description   *string
	Difficulty    *uint `gorm:"check:difficulty >= 1 AND difficulty <= 10"`
	TotalCalories *uint
	TotalPrepTime *string
	Ingredients   []SuggestedIngredient
	Directions    []SuggestedDirection
	accepted      bool
}

type SuggestedIngredient struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"not null;default:null"`
	Quantity          string
	MeasurementUnit   *string
	SuggestedRecipeID uint
}

type SuggestedDirection struct {
	gorm.Model
	ID                uint `gorm:"primaryKey"`
	Description       string
	Time              *uint
	Order             uint `gorm:"not null;default:null"`
	SuggestedRecipeID uint
}
