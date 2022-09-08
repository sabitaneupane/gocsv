package gocsv

import (
	"testing"
)

func TestWriter_Writer(t *testing.T) {
	cases := []struct {
		desc        string
		fileName    string
		hasHeader   bool
		data        CSVData
		expectedErr error
	}{
		{
			"Return exact text",
			"../write.csv",
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
						`=HYPERLINK("http://nsabita.com.np/", "View More")`,
						"59",
					},
				},
			},
			nil,
		},
	}
	for _, tc := range cases {
		actualErr := Writer(tc.fileName, tc.data, tc.hasHeader)
		if actualErr != tc.expectedErr {
			t.Errorf("%s: expected: %s got: %s \n", tc.desc, actualErr, tc.expectedErr)
		}
	}
}

func TestWriter_StringMapper(t *testing.T) {

	cases := []struct {
		desc      string
		hasHeader bool
		data      CSVData
		expected  string
	}{
		{
			"When text contain csv injectable char like =,+,-,@,\t",
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
						`=HYPERLINK("http://nsabita.com.np/", "View More")`,
						"59",
					},
				},
			},
			`"firstname","lastname","age"
"'=cmd|' /C calc'!A0","Doe","23"
"John","'=HYPERLINK(\"\"http://nsabita.com.np/\"\", \"\"View More\"\")","59"
`,
		},
		{
			"When text contain doesnot start with any csv injectable char like =,+,-,@,\t",
			true,
			CSVData{
				Headers: []string{"firstname", "lastname", "age"},
				Body: [][]string{
					{
						"Hari",
						"Doe",
						"23",
					},
					{
						"John",
						`Sharma`,
						"59",
					},
				},
			},
			`"firstname","lastname","age"
"Hari","Doe","23"
"John","Sharma","59"
`,
		},
		{
			"When text contain doesnot start with any csv injectable char like =,+,-,@,\t",
			false,
			CSVData{
				Headers: []string{},
				Body: [][]string{
					{
						"Hari",
						"Doe",
						"23",
					},
					{
						"John",
						`Sharma`,
						"59",
					},
				},
			},
			`"Hari","Doe","23"
"John","Sharma","59"
`,
		},
	}

	for _, tc := range cases {
		got := StringMapper(tc.data, tc.hasHeader)
		ExpectEqual(t, got, tc.expected)
	}
}

func ExpectEqual(t *testing.T, got, expected interface{}) bool {
	if got != expected {
		t.Helper()
		t.Errorf("\n Expected: \n%v \n\n Got: \n%v", expected, got)
		return false
	}
	return true
}
