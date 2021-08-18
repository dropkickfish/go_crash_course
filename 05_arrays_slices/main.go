package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	//Declare and assign
	vegArr := [2]string{"Carrot", "Potato"} // will error if not used - different to fruit Array because fruit Array is then assigned values

	fmt.Println(vegArr)
	fmt.Println(vegArr[1])

	chocSlice := []string{"Cadbury's", "Galaxy", "Mars", "Lindt"} // a slice is the same as an array but witout a predefined length

	fmt.Println(len(chocSlice)) //len finds the length of a slice or array
	fmt.Println(chocSlice[1:3]) //to show a range of items from the array, use this syntax. In this example, it starts at 1 but ends before index 3, so only shows 2 items

}
