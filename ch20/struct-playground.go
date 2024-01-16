package main

import "fmt"

// Define a struct called Person with related fields that represent metadata
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Associate GetFullName method with Person struct to get full name
func (p *Person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	// Create struct instances from the Person struct by declaring variable person1
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Access individual fields from instance of person1 struct
	fmt.Println("Access individual fields from person1 struct:")
	fmt.Println(person1.FirstName)
	fmt.Println(person1.LastName)
	fmt.Println(person1.Age)

	// Call GetFullName method associated with Person struct on person1 struct instance
	fullName := person1.GetFullName()
	fmt.Println("Full name is: ", fullName)
}
