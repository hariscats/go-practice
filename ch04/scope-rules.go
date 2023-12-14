package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" // global scoped era variable

func main() {
	year := 2018 // function scoped year variable

	switch month := rand.Intn(12) + 1; month { // switch block scoped variable
	case 2:
		day := rand.Intn(28) + 1 // each day variable is scoped in each case block
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}
