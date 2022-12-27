package main

import "fmt"

func main() {

	// Creates the function and executes after it
	msg := func(text string) string {
		return fmt.Sprintf("Hello %s", text)
	}("Wilton")

	fmt.Println(msg)
}
