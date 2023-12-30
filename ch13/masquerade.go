package main

import "fmt"

// Assign an anonymous function to variable f at package scope
var f = func() {
	fmt.Println("Dress up for the masquarade")
}

func main() {
	f() // Call the anonymous function by specifying variable name and then parenthesis
}
