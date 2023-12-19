package main

import "fmt"

func main() {
	year := 2023

	// the %T format verb displays a variables type
	// Rather than specifying 'year' twice, %[1] tells to use 1st arg
	fmt.Printf("Type %T for %[1]v\n", year)
}
