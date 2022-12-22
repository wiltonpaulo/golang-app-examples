package main

import "fmt"

func main() {
	var var1 string = "Variable 1"
	var2 := "Variable 2"
	fmt.Println(var1)
	fmt.Println(var2)
	var (
		var3 string = "hahahah"
		var4 string = "hahahah2"
	)
	var5, var6 := "var5", "var6"
	fmt.Println(var3, var4, var5, var6)

	// same as var but you can't change the value
	const const1 string = "Constant 1"
	fmt.Println(const1)
}
