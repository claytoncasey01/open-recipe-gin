package parsers

import (
	"testing"

	"github.com/claytoncasey01/open-recipe-gin/dto"
	"github.com/stretchr/testify/assert"
)

func TestParse_ValidCSV(t *testing.T) {
	parser := CsvParser{}

	tests := []struct {
		name     string
		content  string
		expected *dto.SuggestedRecipeDTO
		wantErr  bool
	}{
		{
			name: "valid CSV",
			content: `NAME,YIELD,TOTAL TIME,TOTAL CALORIES,DIFFICULTY,INGREDIENTS,PREPERATION
LASAGNA,12,90 MINUTES,4530,6,,
,,,,,"1 teaspoon Ingredient 1","1        Mix all ingredients"
,,,,,"2 cups Ingredient 2","2 			Bake for 20 mins"
,,,,,"1 lb Ingredient 3",
`,
			expected: &dto.SuggestedRecipeDTO{
				Name:          "LASAGNA",
				Description:   nil,
				TotalPrepTime: stringPtr("90 MINUTES"),
				TotalCalories: uintPtr(4530),
				Difficulty:    uintPtr(6),
				Ingredients: []dto.SuggestedIngredientDTO{
					{Name: "Ingredient 1", Quantity: "1", MeasurementUnit: stringPtr("teaspoon")},
					{Name: "Ingredient 2", Quantity: "2", MeasurementUnit: stringPtr("cups")},
					{Name: "Ingredient 3", Quantity: "1", MeasurementUnit: stringPtr("lb")},
				},
				Directions: []dto.SuggestedDirectionDTO{
					{Order: 1, Description: "Mix all ingredients"},
					{Order: 2, Description: "Bake for 20 mins"},
				},
			},
			wantErr: false,
		},
		{
			name:     "empty CSV",
			content:  ``,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse(tt.content)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}

func TestParseUnit(t *testing.T) {
	tests := []struct {
		input    string
		expected *uint
	}{
		{"123", uintPtr(123)},
		{"", nil},
		{"abc", nil},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, ParseUint(tt.input))
	}
}

func TestParseQuantity(t *testing.T) {
	tests := []struct {
		input        string
		expectedQty  string
		expectedUnit *string
	}{
		{"1 cup sugar", "1", stringPtr("cup")},
		{"2 tbsp flour", "2", stringPtr("tbsp")},
		{"", "", nil},
	}

	for _, tt := range tests {
		qty, unit := ParseQuantity(tt.input)
		assert.Equal(t, tt.expectedQty, qty)
		assert.Equal(t, tt.expectedUnit, unit)
	}
}

func TestParseName(t *testing.T) {
	tests := []struct {
		ingredient string
		quantity   string
		unit       *string
		expected   string
	}{
		{"1 cup sugar", "1", stringPtr("cup"), "sugar"},
		{"2 tbsp flour", "2", stringPtr("tbsp"), "flour"},
		{"2 cups of milk", "2", stringPtr("cups"), "of milk"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, ParseName(tt.ingredient, tt.quantity, tt.unit))
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func uintPtr(u uint) *uint {
	return &u
}
