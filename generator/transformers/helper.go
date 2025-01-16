package transformers

import (
	"strings"
	"unicode"
)

func ToCamelCase(input string) string {
	isToUpper := true
	var result strings.Builder

	for _, r := range input {
		if r == '_' || r == ' ' {
			isToUpper = true
			continue
		}
		if isToUpper {
			result.WriteRune(unicode.ToUpper(r))
			isToUpper = false
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
