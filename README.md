# gocsv
The gocsv package aims to provide easy CSV 
- Reader
- Writer
- Sanitizer
- Unsanitizer
- StringMapper


[![goreport](https://goreportcard.com/badge/github.com/sabitaneupane/gocsv)](https://goreportcard.com/report/github.com/sabitaneupane/gocsv)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=sabitaneupane_gocsv&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=sabitaneupane_gocsv)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=sabitaneupane_gocsv&metric=bugs)](https://sonarcloud.io/summary/new_code?id=sabitaneupane_gocsv)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=sabitaneupane_gocsv&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=sabitaneupane_gocsv)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=sabitaneupane_gocsv&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=sabitaneupane_gocsv)


## Installation

```
go install github.com/sabitaneupane/gocsv@latest
```

## Requirement
Go 1.18+

### Writer
```
type CSVData struct {
	Headers []string   `json:"headers"`
	Body    [][]string `json:"body"`
}
```

```
func Writer(filename string, data CSVData, hasHeader bool) (err error) {}
```


### Reader
```
type CSVData struct {
	Headers []string   `json:"headers"`
	Body    [][]string `json:"body"`
}
```

```
func Reader(filename string, hasHeader bool) (data CSVData, err error) {}
```

### Sanitizer
```
func Sanitizer(text string) (sanitizedText string) {}
```

### Unsanitizer
```
func Unsanitizer(text string) (unSanitizedText string) {}
```

### StringMapper
```
type CSVData struct {
	Headers []string   `json:"headers"`
	Body    [][]string `json:"body"`
}
```

```
func StringMapper(data CSVData, hasHeader bool) (content string) {}
```