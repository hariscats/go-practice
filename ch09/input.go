package main

import "fmt"

func main() {
	var input string = "yes"

	switch input {
	case "true", "yes", "1":
		fmt.Println(true)
	case "false", "no", "0":
		fmt.Println(false)
	default:
		fmt.Println("Invalid input")
	}
}
