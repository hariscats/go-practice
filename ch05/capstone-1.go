package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Spaceline         Days    Trip type    Price")
	fmt.Println("============================================")

	spaceline := []string{"Virgin Atlantic", "SpaceX", "Space Adventures"}
	tripType := []string{"one-way", "round-trip"}
	secondsInDay := 86400
	kmToMars := 62100000

	numIterations := 10

	for i := 0; i < numIterations; i++ {
		Days := 0

		// assign random index for spaceline
		randomSpaceline := rand.Intn(len(spaceline))
		selectedSpaceline := spaceline[randomSpaceline]

		// assign random index for trip type
		randomTripType := rand.Intn(len(tripType))

		// random index for speed for spaceship between 16 to 30 km/s
		shipSpeed := rand.Intn(16) + 14
		// Choose random speed
		selectShipSpeed := shipSpeed

		if shipSpeed < 24 {
			Days = (kmToMars / shipSpeed) / secondsInDay

		}
	}

}
