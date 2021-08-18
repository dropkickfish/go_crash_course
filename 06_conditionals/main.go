package main

import "fmt"

func main() {
	x := 5
	y := 10

	//if else - if is the same construction, simply without the else block
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y) //%d is a base 10 integer (placeholder)
	} else {
		fmt.Printf("%d is less than or equal to %d\n", y, x) //%d is a base 10 integer (placeholder)
	}

	//else if
	colour := "red"
	if colour == "red" {
		fmt.Println("Colour is red")
	} else if colour == "blue" {
		fmt.Println("Colour is blue")
	} else {
		fmt.Println("Colour is not blue or red")
	}

	//switch (cases)
	switch colour {
	case "red":
		fmt.Println("Colour is red")
	case "blue":
		fmt.Println("Colour is blue")
	default:
		fmt.Println("Colour is neither red nor blue")
	}
}
