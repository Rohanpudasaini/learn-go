package main

import (
	"fmt"

	"github.com/Rohanpudasaini/learn-go/utils"
)

func main() {
	// var a, b int = 20, 0
	utils := utils.Calculator{A: 20, B: 10}
	fmt.Println("Add: ", utils.Add())
	fmt.Println("Subtract: ", utils.Subtract())
	fmt.Println("Multiply: ", utils.Multiply())
	var division, error = utils.Divide()
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Divide: ", division)
	}
}
