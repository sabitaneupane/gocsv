package gocsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Reader(filename string, hasHeader bool) (data CSVData, err error) {
	// Open the file
	recordFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("unable to read input file "+filename, err)
		return
	}
	defer recordFile.Close()

	// Initialize the reader
	csvReader := csv.NewReader(recordFile)

	// Read all the records
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("unable to parse file as CSV for "+filename, err)
		return
	}

	if len(records) == 0 {
		fmt.Println("no CSV records found"+filename, err)
		err = ErrEmptyCSVFile
		return
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("an error encountered while closing file::", err)
		return
	}

	data, err = formatCSVReadData(records, hasHeader)

	if err != nil {
		fmt.Println("an error encountered", err)
	}

	return
}

func formatCSVReadData(records [][]string, hasHeader bool) (newData CSVData, err error) {

	// Using the records content
	if hasHeader {
		newData = CSVData{
			Headers: records[0],
			Body:    records[1:],
		}
	} else {
		newData = CSVData{
			Body: records,
		}
	}

	var headers []string
	if hasHeader {
		for _, h := range newData.Headers {
			header := Unsanitizer(h)
			headers = append(headers, header)
		}
	}

	var body [][]string

	for _, items := range newData.Body {
		var val string
		var row []string

		for _, v := range items {
			val = Unsanitizer(v)
			row = append(row, val)

		}
		body = append(body, row)
	}

	newData = CSVData{
		Headers: headers,
		Body:    body,
	}

	return
}
