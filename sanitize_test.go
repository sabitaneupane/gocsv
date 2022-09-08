package gocsv

import "testing"

func TestCsvSanitizer_Sanitizer(t *testing.T) {
	cases := []struct {
		desc     string
		text     string
		expected string
	}{
		{
			"When text doesnot start with any csv injectable char like =,+,-,@,\t",
			"hello@test.com",
			"\"hello@test.com\"",
		},
		{
			"When text start with csv injectable char like =,+,-,@,\t",
			"=cmd|' /C calc'!A0",
			"\"'=cmd|' /C calc'!A0\"",
		},
		{
			"When text contains double quote",
			`=HYPERLINK("http://nsabita.com.np/", "View More")`,
			`"'=HYPERLINK(\"\"http://nsabita.com.np/\"\", \"\"View More\"\")"`,
		},
	}
	for _, tc := range cases {
		actual := Sanitizer(tc.text)
		if actual != tc.expected {
			t.Errorf("%s: expected: %s got: %s \n", tc.desc, actual, tc.expected)
		}
	}
}
