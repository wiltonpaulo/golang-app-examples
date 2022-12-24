package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	// creates one array with 5 string positions and zero value
	var array1 [5]string
	fmt.Println(array1)

	array1[0] = "First position"
	fmt.Println(array1)

	// array1[5] = "Invalid position"
	// Error: index 5 out of bounds [0:5]

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}
	fmt.Println(array2)

	// Creates the array based on the values when declared - rarely used
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slices with integers. Widely used in golang
	// The number of positions is not fixed, but the datatype is.
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	// Prints the var types of both vars
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
	fmt.Printf("%T\n", slice)
	fmt.Printf("%T\n", array3)

	// Add one more item to an existing slice
	slice = append(slice, 18)
	fmt.Println(slice)

	// Create a slice with a pointer to array2 positions [0, "/1, 2, /3", 4]
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// We are changing the value in the array2 and the slice2 will be modified
	// Because it is a pointer
	array2[1] = "changed Position 2"
	fmt.Println(slice2)

	// Internal arrays
	fmt.Println("--------------------------------")
	fmt.Println("Internal Arrays")
	slice3 := make([]float32, 10, 11) // builtin function to create the array. Eg: make([]type, min_size, max_size)
	fmt.Println(slice3)
	fmt.Printf("lenght: %d capacity: %d \n", len(slice3), cap(slice3))

	// Add more values to extend the slice capacity
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Printf("lenght: %d capacity: %d \n", len(slice3), cap(slice3))

	// Declare a slice with fixed number len/cap and add one value to extend it
	slice4 := make([]float64, 5)
	fmt.Println(slice4)
	fmt.Printf("lenght: %d capacity: %d \n", len(slice4), cap(slice4))
	slice4 = append(slice4, 6)
	fmt.Printf("lenght: %d capacity: %d \n", len(slice4), cap(slice4))
}
