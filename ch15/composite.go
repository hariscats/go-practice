package main

import "fmt"

func main() {
	// composite literal for declaring and initializing an array
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// The ellipsis (...) is a special token that tells the compiler to count elements
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

	// The built-in len() function returns the length of an array
	fmt.Println(len(dwarfs))
	fmt.Println(len(planets))
}
