package main

import "fmt"

type fahrenheit float64
type celsius float64

// celsius converts °F to °C
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var degrees fahrenheit = 45

	fmt.Println(degrees.celsius())
}
