package main

import "fmt"

func main() {
	// Anonymous functions are useful when you need to create functions on-the-fly
	func() {
		fmt.Println("Functions anonymous")
	}() // Declare and call an anonymous function w/o even assigning to variable
}
