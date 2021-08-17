package main

import "fmt"

func main() {
	//Using var
	var name = "Andy"
	var age = 30
	//Using const
	const isCool = true
	//Shorthand
	secondName := "Thackray"        //can only be used in a func
	sport, drink := "snowboard", 30 //2 at once

	fmt.Println(name, secondName, age, isCool, sport, drink) //unused vars throw errors
	fmt.Printf("%T", name)
}
