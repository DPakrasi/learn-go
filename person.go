package main

import "fmt"

type ContactInfo struct {
	zipCode int
	email   string
}

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func (person Person) print() {
	fmt.Printf("%+v", person)
}

func (person *Person) updateFirstName(newFirstName string) {
	(*person).firstName = newFirstName
}
