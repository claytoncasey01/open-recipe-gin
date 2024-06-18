package repositories

import (
	"github.com/claytoncasey01/open-recipe-gin/models"
	"gorm.io/gorm"
)

type SuggestedRecipeRepository interface {
	FindById(id uint) (*models.SuggestedRecipe, error)
	Create(recipe models.SuggestedRecipe) (uint, error)
}

type suggestedRecipeRepository struct {
	DB *gorm.DB
}

func NewSuggestedRecipeRepository(DB *gorm.DB) SuggestedRecipeRepository {
	return &suggestedRecipeRepository{DB}
}

func (r *suggestedRecipeRepository) FindById(id uint) (*models.SuggestedRecipe, error) {
	var suggestedRecipe models.SuggestedRecipe
	if err := r.DB.Preload("Ingredients").Preload("Directions").First(&suggestedRecipe, id).Error; err != nil {
		return nil, err
	}
	return &suggestedRecipe, nil
}

func (r *suggestedRecipeRepository) Create(suggestedRecipe models.SuggestedRecipe) (uint, error) {
	if err := r.DB.Create(&suggestedRecipe).Error; err != nil {
		return 0, err
	}
	return suggestedRecipe.ID, nil
}
