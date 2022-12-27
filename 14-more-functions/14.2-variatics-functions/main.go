package main

import "fmt"

// it uses a range as input value
// - you can only use one range per function
// - it's mandatory to set as the last input item
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}
