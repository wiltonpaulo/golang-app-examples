package main

import "fmt"

func function1() {
	fmt.Println("Running function 1")
}

func function2() {
	fmt.Println("Running function 2")
}

func studentPassed(n1, n2 float32) bool {
	fmt.Println("Checking if the student has passed")
	// defer sends the execution to the end of the function before returning values
	defer fmt.Println("The grade was retrieved and the result was calculated")
	grade := (n1 + n2) / 2
	if grade >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Defer")
	// defer = postpone
	defer function1()
	function2()
	fmt.Println(studentPassed(7, 4))
}
