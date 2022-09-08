package gocsv

import (
	"fmt"
	"testing"
)

func TestWriter_Reader(t *testing.T) {
	cases := []struct {
		desc      string
		fileName  string
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
		actual, _ := Reader(tc.fileName, tc.hasHeader)

		if checkAsStrings(actual, tc.expected) {
			t.Errorf("%s: expected: %v got: %s \n", tc.desc, tc.expected, actual)
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
	}

	for _, tc := range cases {
		actual, _ := formatCSVReadData(tc.data, tc.hasHeader)

		if checkAsStrings(actual, tc.expected) {
			t.Errorf("%s: expected: %v got: %s \n", tc.desc, tc.expected, actual)
		}
	}
}
