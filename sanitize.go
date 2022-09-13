package gocsv

import (
	"fmt"
	"strconv"
	"strings"
)

func Sanitizer(text string) (sanitizedText string) {
	// sanitizing: each double quote in the input should be escaped with another double quote
	sanitizedText = strings.ReplaceAll(text, `"`, `""`)

	// sanitizing: each field containing csvInjectionChar as starting character should start with a single quote
	for _, char := range CsvInjectionChar {
		if strings.HasPrefix(sanitizedText, char) {
			sanitizedText = fmt.Sprintf("'%s", sanitizedText)
			break
		}
	}

	// sanitizing: each input should be wrapped in double quotes
	sanitizedText = strconv.Quote(sanitizedText)

	return sanitizedText
}
