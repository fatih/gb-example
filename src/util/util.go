package util

import (
	"strings"

	"github.com/fatih/camelcase"
)

// SnakeCase is an example exported function that converts a CamelCase string
// to Snake case
func SnakeCase(s string) string {
	return strings.Join(camelcase.Split(s), "_")
}
