package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func ModifyPerson(p Person) {
	p.FirstName = "Alice"
	// Changes to 'p' do not affect the original struct
	// When passing struct as parameter to function, it's passed by value by default.
	// If function needs to modify underlying struct, pass by point or reference
}

func main() {
	person := Person{"Bob", "Smith"}
	ModifyPerson(person)
	fmt.Println(person.FirstName) // Outputs: Bob
}
