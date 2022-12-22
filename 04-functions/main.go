package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalculator(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}

func main() {
	// sum values
	sum := sum(10, 20)
	fmt.Println(sum)

	// function using in a local var
	var f = func(txt string) string {
		fmt.Println(txt, "internal")
		return txt + " return"
	}
	result := f("Text of the function 1")
	fmt.Println(result)

	// mathCalculator function
	sumResult, subtractionResult := mathCalculator(10, 15)
	fmt.Println(sumResult, subtractionResult)

}
