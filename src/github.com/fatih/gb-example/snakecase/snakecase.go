package snakecase

import (
	"strings"

	"github.com/fatih/camelcase"
)

// Convert is an example exported function that converts a CamelCase string
// to Snake case
func Convert(s string) string {
	return strings.Join(camelcase.Split(s), "_")
}
