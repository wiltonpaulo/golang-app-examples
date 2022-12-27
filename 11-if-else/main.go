package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 15

	if number > 15 {
		fmt.Println("More than 15")
	} else if number == 15 {
		fmt.Println("Equals to 15")
	} else {
		fmt.Println("Less than 15")
	}

	// Creates an if statement, a new internal var and evaluate it
	// The variable works only inside the if / it's an internal var
	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Number is more than 0")
	}
}
