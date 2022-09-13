package gocsv

import (
	"fmt"
	"testing"
)

func TestWriter_Reader(t *testing.T) {
	var expectedData CSVData = CSVData{
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
	}

	t.Log("When csv read success")
	{
		t.Run("returns exact text", func(t *testing.T) {
			var filename string = "./example/sample.csv"
			var hasHeader bool = true

			actual, _ := Reader(filename, hasHeader)

			if checkAsStrings(actual, expectedData) {
				t.Errorf("expected: %v got: %s \n", expectedData, actual)
			}
		})
	}

	t.Log("When csv read fails")
	{
		t.Run("returns exact text", func(t *testing.T) {
			var filename string = "./example/empty.csv"
			var hasHeader bool = true
			var expectedErr error = ErrEmptyCSVFile

			_, actualErr := Reader(filename, hasHeader)

			if actualErr != expectedErr {
				t.Errorf("expected: %s got: %s \n", actualErr, expectedErr)
			}
		})
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
