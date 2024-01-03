package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// Accessing an array element outside of bounds leads to a runtime error;
	// uncomment the following line to see it in action:
	// fmt.Println(dwarfs[5])

	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)

	}
	fmt.Println("")
	// The range clause provides a concise and idiomatic way to iterate over an array.
	// It returns both the index and the value of each element.
	fmt.Println("With the range keyword: ")
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}
