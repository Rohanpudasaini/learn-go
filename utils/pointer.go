package utils

import "fmt"

func Something() { fmt.Println("Hello") }

func ChangeValue(val *int) {
	fmt.Println(val) // Prints the memory address
	*val += 10       // Dereferencing pointer and changing the value
}
