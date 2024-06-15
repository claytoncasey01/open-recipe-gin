package routes

import (
	"github.com/claytoncasey01/open-recipe-gin/controllers"
	"github.com/claytoncasey01/open-recipe-gin/repositories"
	"github.com/claytoncasey01/open-recipe-gin/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	recipeRepository := repositories.NewRecipeRepository(db)
	recipeService := services.NewRecipeService(recipeRepository)
	recipeController := controllers.NewRecipeController(recipeService)

	api := router.Group("/api/v1")
	{
		api.GET("/recipes", recipeController.GetAllRecipes)
		api.GET("/recipes/:id", recipeController.GetRecipeById)
		api.POST("/recipes", recipeController.CreateRecipe)
	}
}
