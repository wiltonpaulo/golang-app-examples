package main

import (
	"fmt"
)

func mathCalculation(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	sum, subtraction := mathCalculation(7, 2)
	fmt.Println("sum:", sum, "subtraction:", subtraction)
}
