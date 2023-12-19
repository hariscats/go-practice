package main

import "fmt"

func main() {
	// uint goes from 0 to 255
	var red uint = 255
	red++ // Adding one makes it wrap
	fmt.Println(red)

	// int8 goes up to 127
	var number int8 = 127
	number++ // Add one makes it wrap as well
	fmt.Println(number)

}
