package main

import "fmt"

// coordinate in degrees, minutes, seconds in N/S/E/W hemisphere
type coordinate struct {
	d, m, s float64
	h       rune
}

// decimal converts a DMS (degrees, minutes, seconds) coordinate to decimal degrees
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	// Bradbury Landing: 4°35'22.2" S, 137°26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
}
