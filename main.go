package main

import (
	"log"

	"github.com/claytoncasey01/open-recipe-gin/config"
	"github.com/claytoncasey01/open-recipe-gin/controllers"
	"github.com/claytoncasey01/open-recipe-gin/database"
	"github.com/claytoncasey01/open-recipe-gin/repositories"
	"github.com/claytoncasey01/open-recipe-gin/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db := database.ConnectDB()
	recipeRepository := repositories.NewRecipeRepository(db)
	recipeService := services.NewRecipeService(recipeRepository)
	recipeController := controllers.NewRecipeController(recipeService)

	r := gin.Default()
	r.POST("/recipes", recipeController.CreateRecipe)
	r.GET("/recipes", recipeController.GetAllRecipes)
	r.GET("/recipes/:id", recipeController.GetRecipeById)

	r.Run() // listen and serve on 0.0.0.0:8080
	log.Println("Server started on :8080")
}
