package repositories

import (
	"github.com/claytoncasey01/open-recipe-gin/dto"
	"github.com/claytoncasey01/open-recipe-gin/models"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	FindAll(filters dto.RecipeFilters) ([]models.Recipe, error)
	FindById(id uint) (models.Recipe, error)
	Create(recipe models.Recipe) (uint, error)
}

type recipeRepository struct {
	DB *gorm.DB
}

func NewRecipeRepository(DB *gorm.DB) RecipeRepository {
	return &recipeRepository{DB}
}

func (r *recipeRepository) FindAll(filters dto.RecipeFilters) ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := r.DB

	if filters.Name != "" {
		query = query.Where("name LIKE ?", "%"+filters.Name+"%")
	}
	if filters.Difficulty != "" {
		query = query.Where("difficulty = ?", filters.Difficulty)
	}
	if filters.PrepTime > 0 {
		query = query.Where("total_prep_time <= ?", filters.PrepTime)
	}
	if err := query.Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *recipeRepository) FindById(id uint) (models.Recipe, error) {
	var recipe models.Recipe
	if err := r.DB.First(&recipe, id).Error; err != nil {
		return models.Recipe{}, err
	}
	return recipe, nil
}

func (r *recipeRepository) Create(recipe models.Recipe) (uint, error) {
	if err := r.DB.Create(&recipe).Error; err != nil {
		return 0, err
	}
	return recipe.ID, nil
}
