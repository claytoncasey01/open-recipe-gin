package repositories

import (
	"github.com/claytoncasey01/open-recipe-gin/models"
	"gorm.io/gorm"
)

type SuggestedRecipeRepository interface {
	FindById(id uint) (models.SuggestedRecipe, error)
	Create(recipe models.SuggestedRecipe) (uint, error)
}

type suggestedRecipeRepository struct {
	DB *gorm.DB
}
