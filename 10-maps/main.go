package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// declare a map
	// map[key_type]value_type *fixed types
	user := map[string]string{
		"name":    "John",
		"surname": "Doe",
	}
	fmt.Println(user)
	fmt.Println(user["name"])

	// Nested map and examples to remove items
	user2 := map[string]map[string]string{
		"name": {
			"firstname": "John",
			"surname":   "Doe",
		},
		"course": {
			"school":  "Udemy",
			"subject": "Programing Golang",
		},
	}
	fmt.Println(user2)
	delete(user2["name"], "surname")
	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)

	user2["billing"] = map[string]string{
		"month": "December",
		"email": "johndoe@personalmail.com",
	}
	fmt.Println(user2)
}
