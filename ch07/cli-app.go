package main

import (
	"flag"
	"fmt"
)

func main() {
	var myFlag = flag.String("myFlag", "defaultValue", "Description of flag")

	flag.Parse()

	fmt.Println("Value of myFlag:", *myFlag)
}
