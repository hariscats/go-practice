package main

import "fmt"

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	fahrenheit := (c * 9.0 / 5.0) + 32.0
	return fahrenheit
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(kelvin)
	fmt.Println(kelvin, "° K is ", celsius, "° C")
	fmt.Println(celsius, "° C", fahrenheit, "° F")
}
