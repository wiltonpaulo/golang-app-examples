package main

// function(var *int) {}
// main() { myvar := n function(&myvar) }

import "fmt"

func signalInvertion(number int) int {
	return number * -1
}
func signalInvertionWithPointer(number *int) {
	*number = *number * -1
	// You don't need to return because the change happens direct in the mem address
}

func main() {
	number := 20
	invertedNumber := signalInvertion(number)
	fmt.Println(invertedNumber)
	fmt.Println(number)

	newNumber := 40
	fmt.Println(newNumber)
	// To access the pointer, it needs to add the signal & when call the function
	signalInvertionWithPointer(&newNumber)
	fmt.Println(newNumber)
}
