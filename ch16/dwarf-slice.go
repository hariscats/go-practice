package main

import "fmt"

func main() {
	// Declare and initialize array
	dwarfsArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfsArray)
	fmt.Printf("array %T\n", dwarfsArray)

	// Declare and initialize a slice; no value in square brackets unlike array
	dwarfsSlice := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfsSlice)
	fmt.Printf("slice %T\n", dwarfsSlice)
}
