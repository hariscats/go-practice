// How long does it take to get to Mars
package main

import "fmt"

func main() {
	// can't be changed when trying to assign new value
	const lightSpeed = 299792 // km/s

	// variables must be declared before used
	var distance = 56000000 // km where Earth and Mars are near in their orbits

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000 // Distance where Earth and Mars are far apart
	fmt.Println(distance/lightSpeed, "seconds")
}
