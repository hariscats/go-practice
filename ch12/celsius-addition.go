package main

import "fmt"

func main() {
	type celsius float64

	const degrees = 20
	var temperature celsius = degrees
	fmt.Println(temperature)

	temperature += 10
	fmt.Println(temperature)
}
