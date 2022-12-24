package main

import "fmt"

// There's no native inheritance in golang
// That's the best to way to do it

type person struct {
	name     string
	lastname string
	age      uint8
}

type student struct {
	person
	course string
	school string
}

func main() {
	fmt.Println("Inheritance")

	p1 := person{"Peter", "Parker", 20}
	fmt.Println(p1)

	e1 := student{person{"Mary", "Jane", 19}, "Golang", "Udemy"}
	fmt.Println(e1)
	fmt.Println(e1.name)
	fmt.Println(e1.course)
}
