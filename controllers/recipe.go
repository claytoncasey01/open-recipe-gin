package controllers

import (
	"net/http"
	"strconv"

	"github.com/claytoncasey01/open-recipe-gin/dto"
	"github.com/claytoncasey01/open-recipe-gin/models"
	"github.com/claytoncasey01/open-recipe-gin/services"
	"github.com/gin-gonic/gin"
)

type RecipeController struct {
	service services.RecipeService
}

func NewRecipeController(service services.RecipeService) *RecipeController {
	return &RecipeController{service}
}

func (c *RecipeController) GetAllRecipes(ctx *gin.Context) {
	var filters dto.RecipeFilters

	// Bind query parameters to the filters struct
	if err := ctx.ShouldBindQuery(&filters); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse[[]models.Recipe](nil, err.Error()))
		return
	}

	recipes, err := c.service.GetAllRecipes(filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewResponse[[]models.Recipe](nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dto.NewResponse(&recipes, ""))
}

func (c *RecipeController) GetRecipeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse[[]models.Recipe](nil, "Invalid ID, must be an integer"))
		return
	}

	recipe, err := c.service.GetRecipeById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.NewResponse[[]models.Recipe](nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.NewResponse(&recipe, ""))
}

func (c *RecipeController) CreateRecipe(ctx *gin.Context) {
	var recipe dto.RecipeDTO
	if err := ctx.ShouldBindJSON(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse[dto.RecipeDTO](nil, err.Error()))
		return
	}

	recipeId, err := c.service.CreateRecipe(recipe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewResponse[models.Recipe](nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewResponse(&recipeId, ""))
}
