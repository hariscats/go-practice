package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Constants
	spaceLines := []string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	speedMin, speedMax := 16, 30 // km/s
	priceMin, priceMax := 36, 50 // Million dollars
	distanceToMars := 62100000   // km (as of October 13, 2020)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%-16s %4s %-10s %s\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("==========================================")

	for i := 0; i < 10; i++ {
		spaceLine := spaceLines[rand.Intn(len(spaceLines))]
		speed := rand.Intn(speedMax-speedMin+1) + speedMin // Speed in km/s
		days := distanceToMars / speed / 86400             // Convert seconds to days
		tripType := "One-way"
		if rand.Intn(2) == 1 {
			tripType = "Round-trip"
		}
		price := rand.Intn(priceMax-priceMin+1) + priceMin
		if tripType == "Round-trip" {
			price *= 2
		}

		fmt.Printf("%-16s %4d %-10s $%3d Million\n", spaceLine, days, tripType, price)
	}
}
