package main

import "fmt"

// kelvin is a new type based on float64,
// typically used to represent a temperature value in Kelvin.
type kelvin float64

// sensor function type
type sensor func() kelvin

// realSensor is a function that simulates getting a sensor reading.
// Here it always returns 0 kelvin.
func realSensor() kelvin {
	return 0
}

// calibrate takes a sensor function and an offset value,
// returning a new sensor function that adds the offset to the original sensor's reading.
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		// s() calls the original sensor function,
		// and adds the offset to its result before returning.
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}
