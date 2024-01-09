package main

import (
	"fmt"
)

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")

	// append multiple elements
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")

	fmt.Println(dwarfs)
}
