package parsers

import (
	"mime/multipart"
	"strconv"
	"strings"
)

// Base generic parser interface
type Parser interface {
	Parse(file multipart.File) (any, error)
}

func ParseFile(parser Parser, file multipart.File) {
	parser.Parse(file)
}

// Helpers
func ParseUint(s string) *uint {
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
func ParseQuantity(ingredient string) (string, *string) {
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

func ParseName(ingredient, quantity string, unit *string) string {
	// Remove the quantity and unit from the ingredient string
	name := strings.Replace(ingredient, quantity, "", 1)
	if unit != nil {
		name = strings.Replace(name, *unit, "", 1)
	}

	return strings.TrimSpace(name)
}
