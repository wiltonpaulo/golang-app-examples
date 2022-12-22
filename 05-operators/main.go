package main

import (
	"fmt"
)

func main() {
	// Operator types: + - / * %
	sum := 1 + 2
	subtraction := 1 - 2
	division := 10 / 2
	multiplication := 10 * 5
	remain := 10 % 2
	fmt.Println(sum, subtraction, division, multiplication, remain)

	// relational operators
	fmt.Println(1 > 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// logic operators
	fmt.Println(true && true)
	fmt.Println(1 == 1 && 2 == 1)
	fmt.Println(!(1 == 1 && 2 == 1))

	// unary operators
	number := 10
	number++     // number = number + 1
	number += 15 // number = number + 15
	number--     // number = number - 1
	fmt.Println(number)

	// There's no ternary operator, use if-else instead
	var text string
	if number > 5 {
		text = "Greater than 5"
	} else {
		text = "Less than 5"
	}
	fmt.Println(text)
}
