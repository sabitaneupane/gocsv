package gocsv

import (
	"testing"
)

func TestWriter_Writer(t *testing.T) {
	var data CSVData = CSVData{
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
	}
	t.Log("When csv write success")
	{
		t.Run("returns exact text", func(t *testing.T) {
			var fileName string = "./example/write.csv"
			var hasHeader bool = true
			var expectedErr error = nil

			actualErr := Writer(fileName, data, hasHeader)
			if actualErr != expectedErr {
				t.Errorf("expected: %s got: %s \n", actualErr, expectedErr)
			}
		})
	}

	t.Log("When csv write fails")
	{
		t.Run("returns error", func(t *testing.T) {
			var fileName string = "./sample/write.csv"
			var hasHeader bool = true
			var expectedErr error = nil

			actualErr := Writer(fileName, data, hasHeader)
			if actualErr == expectedErr {
				t.Errorf("expected: %s got: %s \n", actualErr, expectedErr)
			}
		})
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

		if got != tc.expected {
			t.Errorf("%s: expected: %s got: %s \n", tc.desc, tc.expected, got)
		}
	}
}
