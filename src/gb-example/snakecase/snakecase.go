package snakecase

import (
	"strings"

	"github.com/fatih/camelcase"
)

// Convert is an example exported function that converts a CamelCase string
// to Snake case
func Convert(s string) string {
	lower := []string{}
	for _, word := range camelcase.Split(s) {
		lower = append(lower, strings.ToLower(word))
	}

	return strings.Join(lower, "_")
}
