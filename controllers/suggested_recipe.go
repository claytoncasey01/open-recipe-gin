package controllers

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/claytoncasey01/open-recipe-gin/dto"
	"github.com/claytoncasey01/open-recipe-gin/models"
	"github.com/claytoncasey01/open-recipe-gin/parsers"
	"github.com/claytoncasey01/open-recipe-gin/services"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type SuggestedRecipeController struct {
	service services.SuggestedRecipeService
}

func NewSuggestedRecipeController(service services.SuggestedRecipeService) *SuggestedRecipeController {
	return &SuggestedRecipeController{service}
}

// TODO: Implement upload
func (c *SuggestedRecipeController) CreateSuggestedRecipe(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse[dto.SuggestedRecipeDTO](nil, err.Error()))
		return
	}

	fileContent, fileType, err := readAndValidateFile(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewResponse[dto.SuggestedRecipeDTO](nil, err.Error()))
		return
	}

	var suggestedRecipe *dto.SuggestedRecipeDTO

	if fileType == "csv" {
		// Parse the csv
		parser := parsers.CsvParser{}
		suggestedRecipe, err = parsers.ParseFile(parser, fileContent)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.NewResponse[dto.SuggestedRecipeDTO](nil, err.Error()))
			return
		}

	} else {
		apiKey := os.Getenv("OPENAI_API_KEY")
		parser := parsers.OpenAIParser{
			Client: openai.NewClient(apiKey),
		}
		suggestedRecipe, err = parsers.ParseFile(parser, fileContent)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.NewResponse[dto.SuggestedRecipeDTO](nil, err.Error()))
			return
		}

	}
	// Create the recipe and get the id
	suggestedRecipeID, err := c.service.CreateSuggestedRecipe(*suggestedRecipe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewResponse[dto.SuggestedRecipeDTO](nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewResponse(&suggestedRecipeID, ""))
}

func (c *SuggestedRecipeController) GetSuggestedRecipeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse[[]models.SuggestedRecipe](nil, "Invalid ID, must be an integer"))
		return
	}

	recipe, err := c.service.GetSuggestedRecipeById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.NewResponse[[]dto.SuggestedRecipeDTO](nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.NewResponse(&recipe, ""))
}

func (c *SuggestedRecipeController) AcceptSuggestedRecipe(ctx *gin.Context) {
	var recipe dto.SuggestedRecipeDTO
	if err := ctx.ShouldBindJSON(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse[dto.SuggestedRecipeDTO](nil, err.Error()))
		return
	}

	recipeId, err := c.service.AcceptSuggestedRecipe(recipe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewResponse[dto.SuggestedRecipeDTO](nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewResponse(&recipeId, ""))
}

// Helper functions
func readAndValidateFile(file *multipart.FileHeader) (string, string, error) {
	src, err := file.Open()
	var fileType string

	if err != nil {
		return "", "", err
	}
	defer src.Close()

	fileContent, err := io.ReadAll(src)
	if err != nil {
		return "", "", err
	}

	// For now we only support txt and csv files
	if !strings.HasSuffix(file.Filename, ".txt") && !strings.HasSuffix(file.Filename, ".csv") {
		return "", "", errors.New("only txt and csv files are allowed")
	}

	if strings.HasSuffix(file.Filename, ".txt") {
		fileType = "txt"
	} else {
		fileType = "csv"
	}

	return string(fileContent), fileType, nil
}
