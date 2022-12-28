package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Trying to recover the execution!")
	}
}

func studentPassed(n1, n2 float64) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}
	panic("I don't know that to do! Is your grade 6???")
}

func main() {
	fmt.Println("Panic and Recover")
	fmt.Println(studentPassed(7, 5))
}
