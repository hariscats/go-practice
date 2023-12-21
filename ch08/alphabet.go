package main

import (
	"fmt"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	fmt.Println(len(alphabet), "bytes")

	// Count the number of runes in the string
	fmt.Println(len([]rune(alphabet)), "runes")

}
