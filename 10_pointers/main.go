package main

import "fmt"

func main() {
	a := 5
	b := &a // this creates a pointer to the memory address storing a

	fmt.Println(a, b)
	fmt.Printf("a is an %T, b is an %T\n", a, b)

	// use * to read val from address
	fmt.Println(*b)

	// change val with pointer
	*b = 10
	fmt.Println(a)
}
