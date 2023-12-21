package main

import "fmt"

// This implements a simple ROT13 cipher
func main() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13   // Rotate 13 places forward in the alphabet (mod 26)
			if c > 'z' { // If we rotated too far, adjust backwards
				c = c - 26 // by going back 26 places (the length of the alphabet)
			}
		}
		fmt.Printf("%c", c)
	}
}
