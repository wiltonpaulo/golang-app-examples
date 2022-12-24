package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Struct")
	var user01 user
	user01.name = "John"
	user01.age = 27
	fmt.Println(user01.name, user01.age)

	exampleAddress := address{"Example street", 0}

	user02 := user{"Andrew", 32, exampleAddress}
	fmt.Println(user02.name, user02.age)

	user03 := user{name: "David", address: exampleAddress}
	fmt.Println(user03)
}
