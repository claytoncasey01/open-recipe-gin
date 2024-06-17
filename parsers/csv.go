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
		TotalCalories: parseUint(records[0][3]),
		Difficulty:    parseUint(records[0][4]),
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
			quantity, unit := parseQuantity(record[1])
			ingredients = append(ingredients, models.Ingredient{
				Name:            parseName(record[1], quantity, unit),
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

func parseUint(s string) *uint {
	if s == "" {
		return nil
	}
	intVal, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	uintVal := uint(intVal)

	return &uintVal
}

// Handles parsing out the quantity and measurement unit from the ingredient string
func parseQuantity(ingredient string) (string, *string) {
	// Splits the ingredient string into parts based on whitespace
	parts := strings.Fields(ingredient)
	if len(parts) < 2 {
		return "", nil
	}

	quantity := parts[0]
	unit := ""

	for _, part := range parts[1:] {
		for _, q := range commonQuantities {
			// We found the unit, so we break out of the loop
			if strings.Contains(part, q) {
				unit = part
				break
			}
		}
		// We found a unit, no need to continue searching
		if unit != "" {
			break
		}
	}

	if unit == "" {
		return quantity, nil
	}

	return quantity, &unit
}

func parseName(ingredient, quantity string, unit *string) string {
	// Remove the quantity and unit from the ingredient string
	name := strings.Replace(ingredient, quantity, "", 1)
	if unit != nil {
		name = strings.Replace(name, *unit, "", 1)
	}

	return strings.TrimSpace(name)
}
