package main

import "fmt"

func main() {
	// long way
	i := 4
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//short way
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// fizz buzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
