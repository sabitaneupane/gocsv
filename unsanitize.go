package gocsv

import "strings"

func Unsanitizer(text string) (unSanitizedText string) {
	// Unsanitizing: Removing escaping double quote from text value
	unSanitizedText = strings.ReplaceAll(text, `\"\"`, `"`)

	// Unsanitizing: Removing double quotes wrapper around text value
	unSanitizedText = strings.Trim(unSanitizedText, `"`)

	// Unsanitizing: Removing single quote from start of text value
	unSanitizedText = strings.TrimPrefix(unSanitizedText, `'`)

	return unSanitizedText
}
