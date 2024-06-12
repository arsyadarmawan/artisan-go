package pkg

import (
	"strings"
	"unicode"
)

func ToPascalCase(str string) string {
	words := strings.FieldsFunc(str, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsDigit(c)
	})
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}
	return strings.Join(words, "")
}

// ReplacePlaceholders replaces placeholders in a template string with values from a map.
func ReplacePlaceholders(template string, values map[string]string) string {
	for placeholder, value := range values {
		template = strings.ReplaceAll(template, placeholder, value)
	}
	return template
}
