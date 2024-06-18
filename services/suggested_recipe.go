package services

import (
	"github.com/claytoncasey01/open-recipe-gin/dto"
	"github.com/claytoncasey01/open-recipe-gin/repositories"
)

type SuggestedRecipeService interface {
	CreateSuggestedRecipe(recipe dto.SuggestedRecipeDTO) (uint, error)
	GetSuggestedRecipeById(id uint) (*dto.SuggestedRecipeDTO, error)
	AcceptSuggestedRecipe(suggestedRecipe dto.SuggestedRecipeDTO) (uint, error)
}

type suggestedRecipeService struct {
	repo       repositories.SuggestedRecipeRepository
	recipeRepo repositories.RecipeRepository
}

func NewSuggestedRecipeService(repo repositories.SuggestedRecipeRepository, recipeRepo repositories.RecipeRepository) SuggestedRecipeService {
	return &suggestedRecipeService{repo, recipeRepo}
}

func (s *suggestedRecipeService) CreateSuggestedRecipe(recipe dto.SuggestedRecipeDTO) (uint, error) {
	suggestedRecipeModel := dto.SuggestedRecipeModelFromDTO(recipe)

	return s.repo.Create(suggestedRecipeModel)
}

func (s *suggestedRecipeService) GetSuggestedRecipeById(id uint) (*dto.SuggestedRecipeDTO, error) {
	suggestedRecipeModel, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	suggestedRecipeDTO := dto.SuggestedRecipeDTOFromModel(*suggestedRecipeModel)
	return &suggestedRecipeDTO, nil
}

// Creates a Recipe from a SuggestedRecipe
func (s *suggestedRecipeService) AcceptSuggestedRecipe(recipe dto.SuggestedRecipeDTO) (uint, error) {
	recipeModel := dto.RecipeModelFromSuggestedRecipeDTO(recipe)

	return s.recipeRepo.Create(recipeModel)
}
