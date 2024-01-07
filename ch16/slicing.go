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

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	// Index into a slice like an array
	fmt.Println(gasGiants[0])

	giants := planets[4:8]
	// Create a slice from a slice
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice)

	iceGiantsMarkII := iceGiants
	// Assigning a new value to an element of a slice modifies the underlying planets array.
	// The change will be visible through the other slices:
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiants, iceGiantsMarkII, ice)
}
