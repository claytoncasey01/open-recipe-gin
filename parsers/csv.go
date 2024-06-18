package parsers

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

	"github.com/claytoncasey01/open-recipe-gin/dto"
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
func (p CsvParser) Parse(content string) (*dto.SuggestedRecipeDTO, error) {
	csvReader := csv.NewReader(strings.NewReader(content))
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return nil, fmt.Errorf("No records found in CSV file")
	}

	// Parse the recipe details from the second row (first is headers)
	suggestedRecipe := dto.SuggestedRecipeDTO{
		Name:          records[1][0],
		TotalPrepTime: &records[1][2],
		TotalCalories: ParseUint(records[1][3]),
		Difficulty:    ParseUint(records[1][4]),
	}

	// Parse ingredients and directions from the rest of the rows
	var ingredients []dto.SuggestedIngredientDTO
	var directions []dto.SuggestedDirectionDTO

	for _, record := range records[2:] {
		// Skip the rows that are completely empty
		if len(record) < 7 || strings.TrimSpace(record[5]) == "" && strings.TrimSpace(record[6]) == "" {
			continue
		}

		// Ingredient Rows
		if strings.TrimSpace(record[5]) != "" {
			quantity, unit := ParseQuantity(record[5])
			ingredients = append(ingredients, dto.SuggestedIngredientDTO{
				Name:            ParseName(record[5], quantity, unit),
				Quantity:        quantity,
				MeasurementUnit: unit,
			})
		}

		// Direction rows
		if strings.TrimSpace(record[6]) != "" {
			// Remove leading/trailing whitespace and extra spaces
			description := strings.TrimSpace(record[6])
			// Remove leading/trailing tabs
			description = strings.Trim(description, "\t")
			// Remove extra spaces
			description = strings.Join(strings.Fields(description), " ")
			// Check if the description starts with a number
			order, err := strconv.Atoi(strings.Split(description, " ")[0])
			if err != nil {
				return nil, fmt.Errorf("Invalid direction order: %s", record[6])
			}
			// Remove the order from the description
			description = strings.TrimPrefix(description, strconv.Itoa(order))
			// Remove any leftover leading/trailing whitespace and extra spaces after removing the order
			description = strings.TrimSpace(description)
			directions = append(directions, dto.SuggestedDirectionDTO{
				Order:       uint(order),
				Description: description,
			})
		}
	}

	suggestedRecipe.Ingredients = ingredients
	suggestedRecipe.Directions = directions

	return &suggestedRecipe, nil
}
