package parsers

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/claytoncasey01/open-recipe-gin/models"
)

type CsvParser struct{}

var commonQuantities = []string{
	"teaspoon", "teaspoons", "tsp",
	"tablespoon", "tablespoons", "tbsp",
	"cup", "cups",
	"ounce", "ounces", "oz",
	"pound", "pounds", "lb", "lbs",
	"gram", "grams", "g",
	"kilogram", "kilograms", "kg",
	"liter", "liters", "l",
	"milliliter", "milliliters", "ml",
}

// Parses the CSV file and returns a recipe
// This CSV format is as follows:
// Name,Description,TotalPrepTime,TotalCalories,Difficulty, INGREDIENTS, DIRECTIONS
// The first row contains the recipe details
// The rest of the rows contain data for the ingredients and directions
func (p CsvParser) Parse(file multipart.File) (any, error) {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return nil, fmt.Errorf("No records found in CSV file")
	}

	// Parse the recipe details from the first row
	recipe := models.Recipe{
		Name:          records[0][0],
		TotalPrepTime: &records[0][2],
		TotalCalories: ParseUint(records[0][3]),
		Difficulty:    ParseUint(records[0][4]),
	}

	// Parse ingredients and directions from the rest of the rows
	var ingredients []models.Ingredient
	var directions []models.Direction

	for i, record := range records[1:] {
		// Skip the empty columns for name, description, preptime, calories, and difficulty
		if strings.TrimSpace(record[1]) == "" {
			continue
		}
		// Ingredient Rows
		if i < 5 {
			quantity, unit := ParseQuantity(record[1])
			ingredients = append(ingredients, models.Ingredient{
				Name:            ParseName(record[1], quantity, unit),
				Quantity:        quantity,
				MeasurementUnit: unit,
			})
		} else {
			// Direction rows
			order, err := strconv.Atoi(strings.Split(record[1], " ")[0])
			if err != nil {
				return nil, fmt.Errorf("Invalid direction order: %s", record[1])
			}
			directions = append(directions, models.Direction{
				Order:       uint(order),
				Description: record[1],
			})
		}
	}

	recipe.Ingredients = ingredients
	recipe.Directions = directions

	return recipe, nil
}
