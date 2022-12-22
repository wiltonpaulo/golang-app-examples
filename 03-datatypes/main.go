package main

import "fmt"

func main() {
	// int(will use the default arch), int8, int16, int32, int64
	var number int = 100000000000
	fmt.Println(number)

	// unasigned int
	var number2 uint = 100000000000
	fmt.Println(number2)

	// alias int32 = rune
	var number3 rune = 1234
	fmt.Println(number3)

	// byte = unint8
	var number4 byte = 123
	fmt.Println(number4)

	// real float numbers
	var realNumber1 float32 = 123.45
	var realNumber2 float64 = 123.45
	fmt.Println(realNumber1, realNumber2)

	// string
	var str string = "text"
	fmt.Println(str)

	// char ascii number
	char := 'B'
	fmt.Println(char)

	// booleans
	var boolean bool
	fmt.Println(boolean)

	// print types
	fmt.Printf("number: %T\n", number)
	fmt.Printf("realNumber1: %T\n", realNumber1)
	fmt.Printf("str: %T\n", str)
	fmt.Printf("boolean: %T\n", boolean)
}
