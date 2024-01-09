package main

import (
	"fmt"
)

func main() {
	// Avoid extra allocations and copies by preallocating a slice with the built-in make function
	dwarfs := make([]string, 0, 10)

	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")

	fmt.Println(len(dwarfs), cap(dwarfs))
}
