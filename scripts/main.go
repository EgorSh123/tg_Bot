package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("cit.csv")
	if err != nil {
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true

	countries := make(map[string]struct{})
	cities := make(map[string]struct{})

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}

		// Проверяем что есть минимум 3 столбца
		if len(record) < 3 {
			continue
		}

		country := clean(record[0])
		city := clean(record[2])

		if country != "" {
			countries[country] = struct{}{}
		}
		if city != "" {
			cities[city] = struct{}{}
		}
	}

	writeToFile("countries.txt", countries)
	writeToFile("cities.txt", cities)
}

func clean(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, `"`)
	return s
}

func writeToFile(filename string, data map[string]struct{}) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for value := range data {
		writer.WriteString(value + "\n")
	}
}
