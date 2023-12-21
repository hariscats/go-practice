package main

import "fmt"

func main() {
	c := 'a'
	fmt.Printf("%c\n", c) // Prints "a"
	c = c + 3             // Shifts it 3 letters forward
	fmt.Printf("%c\n", c) // Prints "d"
}
