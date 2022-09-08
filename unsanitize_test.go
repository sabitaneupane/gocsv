package gocsv

import "testing"

func TestCsvSanitizer_Unsanitizer(t *testing.T) {
	cases := []struct {
		desc     string
		text     string
		expected string
	}{
		{
			"When text contains normal text",
			"\"hello@test.com\"",
			"hello@test.com",
		},
		{
			"When text contains single quote at the start of text",
			"\"'=cmd|' /C calc'!A0\"",
			"=cmd|' /C calc'!A0",
		},
		{
			"When text contains double quote escaping",
			`"'=HYPERLINK(\"\"http://nsabita.com.np/\"\", \"\"View More\"\")"`,
			`=HYPERLINK("http://nsabita.com.np/", "View More")`,
		},
	}
	for _, tc := range cases {
		actual := Unsanitizer(tc.text)
		if actual != tc.expected {
			t.Errorf("%s: expected: %s got: %s \n", tc.desc, actual, tc.expected)
		}
	}
}
