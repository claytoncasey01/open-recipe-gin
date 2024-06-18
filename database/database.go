package database

import (
	"log"
	"os"

	"github.com/claytoncasey01/open-recipe-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Make sure we run the AutoMigrate on connect to sync any
	// model changes to the DB.
	db.AutoMigrate(&models.Recipe{}, &models.Ingredient{}, &models.Direction{}, &models.SuggestedRecipe{}, &models.SuggestedIngredient{}, &models.SuggestedDirection{})

	return db
}
