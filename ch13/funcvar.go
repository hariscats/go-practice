package main

import "fmt"

func main() {
	// Assign anonymous function with parameters to variable f within function scope
	f := func(message string) {
		fmt.Println(message)
	}
	f("Go to the party") // Pass in argument to variable f function
}
