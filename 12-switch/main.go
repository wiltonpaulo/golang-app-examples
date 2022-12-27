package main

import (
	"fmt"
)

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
}

func shiftOfDay(number int) string {
	switch {
	case number < 12:
		return "Good morning!!!"
		// fallthrough -> sends the condition to the next case
	case number < 18:
		return "Good afternoon!!!"
	case number <= 23:
		return "Good evening!!!"
	default:
		return "Hello!!!"
	}
}

func main() {
	fmt.Println("Switch")
	day := dayOfWeek(3)
	shift := shiftOfDay(10)
	fmt.Printf("Today is %v. %v\n", day, shift)
}
