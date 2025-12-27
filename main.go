package main

import (
	"fmt"

	"github.com/Rohanpudasaini/learn-go/utils"
)

// var c, python, java bool
var c, python, java = true, 12, "no!"

func main() {
	// var a, b int = 20, 0
	i := 12
	f := float64(i)
	u := uint(f)
	fmt.Println(c, python, java, i)
	fmt.Printf("Type %T value: %v \n", f, f)
	fmt.Println(u)
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
