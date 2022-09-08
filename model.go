package gocsv

type CSVData struct {
	Headers []string   `json:"headers"`
	Body    [][]string `json:"body"`
}
