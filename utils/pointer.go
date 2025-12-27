package utils

import "fmt"

func Something() { fmt.Println("Hello") }

func ChangeValue(val *int) {
	fmt.Println(val) // Prints the memory address
	*val += 10       // Dereferencing pointer and changing the value
}

func Swap(val1, val2 *int) {
	temp := *val1
	*val1 = *val2
	*val2 = temp
}
