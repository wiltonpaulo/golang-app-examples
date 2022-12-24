package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var variable1 int = 10
	var variable2 int = variable1

	variable1++

	fmt.Println(variable1, variable2)

	// Pointer is a reference to access the information in memory
	var variable3 int
	var pointer *int

	variable3 = 100
	pointer = &variable3

	// Gets the address of a pointer
	fmt.Println(variable3, pointer)

	variable3 = 150

	// Retrieve the information in the memory address - unreference
	fmt.Println(variable3, *pointer)

}
