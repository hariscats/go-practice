package main

import (
	"fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	// Write a loop to decipher the cipherText given the keyword into plainText
	keywordIndex := 0
	plainText := ""

	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i]
		k := keyword[keywordIndex]

		cipherLetter := c - 'A'
		keyLetter := k - 'A'

		plainLetter := (cipherLetter - keyLetter + 26) % 26

		plainText += string(plainLetter + 'A')

		keywordIndex++
		keywordIndex %= len(keyword)
	}
	fmt.Println(plainText)
}
