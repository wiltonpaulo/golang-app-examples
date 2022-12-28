package main

import "fmt"

func closure() func() {
	text := "Inside the closure function"

	function := func() {
		fmt.Println(text)
	}
	return function
}
func main() {
	fmt.Println("Closure")
	text := "Inside the main function"
	fmt.Println(text)

	// It's like a memory from the source function
	// You can see that the text var came from the function
	newFunction := closure()
	newFunction()
}
