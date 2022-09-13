package gocsv

import (
	"fmt"
	"os"
)

func Writer(filename string, data CSVData, hasHeader bool) (err error) {
	// open the file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("an error encountered when creating file::", err)
		return
	}

	// format the csv data
	newData := StringMapper(data, hasHeader)

	// write all the records
	_, err = f.WriteString(newData)
	if err != nil {
		fmt.Println("an error encountered writing to a file::", err)
		return
	}
	return
}

func StringMapper(data CSVData, hasHeader bool) (content string) {
	content = ""
	separator := ""
	lineBreaker := "\n"

	// checks if csv containes headers and format csvdata to string
	if hasHeader {
		for _, h := range data.Headers {
			header := Sanitizer(h)

			content += separator + header
			separator = ","
		}
		content += lineBreaker
	}

	// format csvdata to string
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
