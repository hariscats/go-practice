package main

import (
	"fmt"
)

// Converts Celsius to Fahrenheit
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

// Converts Fahrenheit to Celsius
func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Draws a table given headers and rows
func drawTable(header1, header2 string, rows [][]string) {
	separator := "======================="
	fmt.Println(separator)
	fmt.Printf("| %-8s | %-8s |\n", header1, header2)
	fmt.Println(separator)
	for _, row := range rows {
		fmt.Printf("| %-8s | %-8s |\n", row[0], row[1])
	}
	fmt.Println(separator)
}

func main() {
	// Generate data for Celsius to Fahrenheit
	var cToFRows [][]string
	for c := -40.0; c <= 100; c += 5 {
		f := celsiusToFahrenheit(c)
		cToFRows = append(cToFRows, []string{fmt.Sprintf("%.1f", c), fmt.Sprintf("%.1f", f)})
	}

	// Generate data for Fahrenheit to Celsius
	var fToCRows [][]string
	for f := -40.0; f <= 212; f += 5 {
		c := fahrenheitToCelsius(f)
		fToCRows = append(fToCRows, []string{fmt.Sprintf("%.1f", f), fmt.Sprintf("%.1f", c)})
	}

	// Draw the tables
	drawTable("°C", "°F", cToFRows)
	drawTable("°F", "°C", fToCRows)
}
