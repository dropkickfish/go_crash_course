package main

import "fmt"

func main() {
	// Define maps
	emails := make(map[string]string)

	// assign kv
	emails["Andy"] = "andy@email.com"
	emails["Bob"] = "bob@email.com"
	emails["Mike"] = "Mike@email.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// delete one
	delete(emails, "Bob")

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Mike"])

	// assign KV on declaration
	roles := map[string]string{"Andy": "Engineer", "Olga": "Support"} //note absence of "make"

	fmt.Println(roles)
	fmt.Println(len(roles))
	fmt.Println(roles["Andy"])
}
