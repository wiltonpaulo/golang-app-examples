package main

import "fmt"

var n int

func init() {
	fmt.Println("Running the init function")
	n = 10
}

func main() {
	fmt.Println("Main function is being executed")
	fmt.Println(n)
}
