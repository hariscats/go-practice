package main

import (
	"fmt"
)

func main() {
	age := 41
	marsAge := float64(age) // type conversion from integer to float64

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays

	fmt.Println("I am", age, "years old, which is", marsAge, "in Martian years.")
}
