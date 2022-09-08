package gocsv

import "strings"

func Unsanitizer(text string) (unSanitizedText string) {
	// unsanitizing: removing escaping double quote from text value
	unSanitizedText = strings.ReplaceAll(text, `\"\"`, `"`)

	// unsanitizing: removing double quotes wrapper around text value
	unSanitizedText = strings.Trim(unSanitizedText, `"`)

	// unsanitizing: removing single quote from start of text value
	unSanitizedText = strings.TrimPrefix(unSanitizedText, `'`)

	return unSanitizedText
}
