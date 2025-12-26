package main

import (
	"fmt"

	"github.com/Rohanpudasaini/learn-go/utils"
)

func main() {
	var a, b int = 20, 0
	fmt.Println("Add: ", utils.Add(a, b))
	fmt.Println("Subtract: ", utils.Subtract(a, b))
	fmt.Println("Multiply: ", utils.Multiply(a, b))
	var division, error = utils.Divide(a, b)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Divide: ", division)
	}
}
