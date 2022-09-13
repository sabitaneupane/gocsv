package gocsv

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func TestWriter_Reader(t *testing.T) {
	cases := []struct {
		desc      string
		filename  string
		hasHeader bool
		expected  CSVData
	}{
		{
			"Return exact text",
			"./example/sample.csv",
			true,
			CSVData{
				Headers: []string{"firstname", "lastname", "age"},
				Body: [][]string{
					{
						"=cmd|' /C calc'!A0",
						"Doe",
						"23",
					},
					{
						"John",
						"Doe",
						"59",
					},
				},
			},
		},
	}
	for _, tc := range cases {

		recordFile, err := os.Open(tc.filename)
		if err != nil {
			t.Errorf("Did not expect error but got %s", err.Error())
		}
		defer recordFile.Close()

		csvReader := csv.NewReader(recordFile)
		records, err := csvReader.ReadAll()
		if err != nil {
			t.Errorf("Did not expect error but got %s", err.Error())
		}

		// verify is csv file is empty
		if len(records) == 0 {
			err = ErrEmptyCSVFile
			t.Errorf("Did not expect error but got %s", err.Error())
			return
		}

		err = recordFile.Close()
		if err != nil {
			t.Errorf("Did not expect error but got %s", err.Error())
		}

		actualData, err := formatCSVReadData(records, tc.hasHeader)

		if err != nil {
			fmt.Println("an error encountered", err)
			return
		}

		if checkAsStrings(actualData, tc.expected) {
			t.Errorf("%s: expected: %v got: %s \n", tc.desc, tc.expected, actualData)
		}
	}
}

func checkAsStrings(a, b interface{}) bool {
	return fmt.Sprintf("%v", a) != fmt.Sprintf("%v", b)
}

func TestWriter_formatCSVReadData(t *testing.T) {
	cases := []struct {
		desc      string
		hasHeader bool
		data      [][]string
		expected  CSVData
	}{
		{
			"Return exact text",
			true,
			[][]string{
				{"firstname", "lastname", "age"},
				{"=cmd|' /C calc'!A0", "Doe", "23"},
				{"John", "Doe", "59"},
			},
			CSVData{
				Headers: []string{"firstname", "lastname", "age"},
				Body: [][]string{
					{
						"=cmd|' /C calc'!A0",
						"Doe",
						"23",
					},
					{
						"John",
						"Doe",
						"59",
					},
				},
			},
		},
		{
			"Return exact text without headers",
			false,
			[][]string{
				{"=cmd|' /C calc'!A0", "Doe", "23"},
				{"John", "Doe", "59"},
			},
			CSVData{
				Headers: []string{},
				Body: [][]string{
					{
						"=cmd|' /C calc'!A0",
						"Doe",
						"23",
					},
					{
						"John",
						"Doe",
						"59",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		actual, _ := formatCSVReadData(tc.data, tc.hasHeader)

		if checkAsStrings(actual, tc.expected) {
			t.Errorf("%s: expected: %v got: %s \n", tc.desc, tc.expected, actual)
		}
	}
}
