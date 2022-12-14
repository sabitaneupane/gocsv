package gocsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Reader(filename string, hasHeader bool) (data CSVData, err error) {
	// open the file
	recordFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("unable to read input file "+filename, err)
		return
	}
	defer recordFile.Close()

	// initialize the reader
	csvReader := csv.NewReader(recordFile)

	// read all the records
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("unable to parse file as CSV for "+filename, err)
		return
	}

	// verify is csv file is empty
	if len(records) == 0 {
		fmt.Println("empty csv file")
		err = ErrEmptyCSVFile
		return
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("an error encountered while closing file::", err)
		return
	}

	data = formatCSVReadData(records, hasHeader)

	return
}

func formatCSVReadData(records [][]string, hasHeader bool) (newData CSVData) {
	// using the records content
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

	//  format csv header to CSVData format
	var headers []string
	if hasHeader {
		for _, h := range newData.Headers {
			header := Unsanitizer(h)
			headers = append(headers, header)
		}
	}

	//  format csv body to CSVData format
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
