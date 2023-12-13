package main

import (
	"fmt"
)

func main() {
	var i interface{}

	// Assign different types of values to 'i' and use the type switch.
	// Here we're using an integer value.
	i = 42
	switch v := i.(type) {
	case int:
		fmt.Printf("i is an integer: %v\n", v)
	case string:
		fmt.Printf("i is a string: %v\n", v)
	default:
		fmt.Printf("i is of a different type: %T\n", v)
	}

	// Now let's assign a string value to 'i' and use the type switch again.
	i = "hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("i is an integer: %v\n", v)
	case string:
		fmt.Printf("i is a string: %v\n", v)
	default:
		fmt.Printf("i is of a different type: %T\n", v)
	}
}
