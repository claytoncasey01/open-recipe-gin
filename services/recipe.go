package services

import (
	"github.com/claytoncasey01/open-recipe-gin/dto"
	"github.com/claytoncasey01/open-recipe-gin/models"
	"github.com/claytoncasey01/open-recipe-gin/repositories"
)

type RecipeService interface {
	GetAllRecipes(filters dto.RecipeFilters) ([]models.Recipe, error)
	GetRecipeById(id uint) (models.Recipe, error)
	CreateRecipe(recipe models.Recipe) (uint, error)
}

type recipeService struct {
	repo repositories.RecipeRepository
}

func NewRecipeService(repo repositories.RecipeRepository) RecipeService {
	return &recipeService{repo}
}

func (s *recipeService) GetAllRecipes(filters dto.RecipeFilters) ([]models.Recipe, error) {
	return s.repo.FindAll(filters)
}

func (s *recipeService) GetRecipeById(id uint) (models.Recipe, error) {
	return s.repo.FindById(id)
}

func (s *recipeService) CreateRecipe(recipe models.Recipe) (uint, error) {
	return s.repo.Create(recipe)
}
