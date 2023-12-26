package main

import "fmt"

func main() {
	type celsius float64 // Declaring a new type called celsius with underlying type float64

	var temperature celsius = 20 // Declaring a variable of type celsius and assigning a value to it (20) of type float64

	fmt.Println(temperature)
}
