package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringPtr(s string) *string {
	return &s
}

func uintPtr(u uint) *uint {
	return &u
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
