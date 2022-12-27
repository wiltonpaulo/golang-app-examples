package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	// Basic for loop with increment
	// i := 0
	// for i < 3 {
	// 	fmt.Println("Increment i", i)
	// 	time.Sleep(time.Second)
	// 	i++
	// }

	// Increment using an internal var in the condition statement
	for j := 0; j < 3; j++ {
		fmt.Println("Increment j", j)
		time.Sleep(time.Second)
	}

	// Iterate for in a list - array or slice
	names := []string{"John", "David", "Lucas"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for index, letter := range "Word" {
		fmt.Println(index, string(letter))
	}

	// Iterate for in a dict, map, hash
	user := map[string]string{
		"name":     "John",
		"lastname": "Doe",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	// Infinite loop
	// for {
	// 	fmt.Println("print anything")
	// }

	// *** It's not possible to iterate a struct
}
