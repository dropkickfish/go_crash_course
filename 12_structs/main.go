package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int // data of same types can all be on same line
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getsMarried method (pointer receiver)
func (p *Person) getsMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Andy", lastName: "Thackray", city: "Rotherham", gender: "M", age: 30}
	fmt.Println(person1)

	// init person using struct without args
	person2 := Person{"Rachel", "Thackray", "Rotherham", "F", 28}
	fmt.Println(person2)

	// print single field
	fmt.Println(person1.firstName)

	// using the value receiver from line 18
	fmt.Println(person1.greet())

	// using a pointer receiver first
	person1.hasBirthday()
	fmt.Println(person1.greet())

	// getsMarried
	person2.getsMarried("Morint")
	fmt.Println(person2.greet())
}
