package gocsv

import (
	"fmt"
	"os"
)

func Writer(filename string, data CSVData, hasHeader bool) (err error) {
	// Open the file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("An error encountered when creating file::", err)
	}

	// Format the csv data
	newData := StringMapper(data, hasHeader)

	// Write all the records
	_, err = f.WriteString(newData) // returns error
	if err != nil {
		fmt.Println("An error encountered writing to a file::", err)
	}
	return
}

func StringMapper(data CSVData, hasHeader bool) (content string) {
	content = ""
	separator := ""
	lineBreaker := "\n"

	if hasHeader {
		for _, h := range data.Headers {
			header := Sanitizer(h)

			content += separator + header
			separator = ","

		}

		content += lineBreaker
	}

	for _, items := range data.Body {
		separator = ""
		var val string

		for _, v := range items {
			val = Sanitizer(v)

			content += separator + val
			separator = ","

		}
		content += lineBreaker

	}

	return
}
