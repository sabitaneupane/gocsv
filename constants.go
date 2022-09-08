package gocsv

import "errors"

var CsvInjectionChar = []string{
	"=",
	"+",
	"-",
	"@",
	"\t",
}

var ErrEmptyCSVFile = errors.New("empty csv file")
