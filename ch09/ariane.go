package main

import "fmt"

func main() {
	var bh float64 = 32767
	var h = int8(bh) // 32767 overflows int8
	fmt.Println(h)
}
