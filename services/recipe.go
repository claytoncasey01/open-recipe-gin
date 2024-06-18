package services

import (
	"github.com/claytoncasey01/open-recipe-gin/dto"
	"github.com/claytoncasey01/open-recipe-gin/repositories"
)

type RecipeService interface {
	GetAllRecipes(filters dto.RecipeFilters) ([]dto.RecipeDTO, error)
	GetRecipeById(id uint) (*dto.RecipeDTO, error)
	CreateRecipe(recipe dto.RecipeDTO) (uint, error)
}

type recipeService struct {
	repo repositories.RecipeRepository
}

func NewRecipeService(repo repositories.RecipeRepository) RecipeService {
	return &recipeService{repo}
}

func (s *recipeService) GetAllRecipes(filters dto.RecipeFilters) ([]dto.RecipeDTO, error) {
	var recipeDTOs []dto.RecipeDTO
	recipeModels, err := s.repo.FindAll(filters)
	if err != nil {
		return nil, err
	}
	for _, recipe := range recipeModels {
		recipeDTOs = append(recipeDTOs, dto.RecipeDTOFromModel(recipe))
	}
	return recipeDTOs, nil
}

func (s *recipeService) GetRecipeById(id uint) (*dto.RecipeDTO, error) {
	recipe, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	recipeDTO := dto.RecipeDTOFromModel(*recipe)
	return &recipeDTO, nil
}

func (s *recipeService) CreateRecipe(recipe dto.RecipeDTO) (uint, error) {
	recipeModel := dto.RecipeModelFromDTO(recipe)
	return s.repo.Create(recipeModel)
}
