package main

import "fmt"

func main() {
	ids := []int{33, 34, 35, 36, 66, 7, 97}

	// loop through IDs
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// loop through IDs without index
	for _, id := range ids { //nb you need that underscore or it'll error (i not used)
		fmt.Printf("ID: %d\n", id)
	}

	// add IDs together
	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// range with emails
	roles := map[string]string{"Andy": "Engineer", "Olga": "Support"}
	for k, v := range roles {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range roles {
		fmt.Printf("%s\n", k)
	}
}
