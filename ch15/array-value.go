package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	// Copies values from planets array (instead of modifying by reference)
	planetsMarkII := planets

	planets[2] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}
