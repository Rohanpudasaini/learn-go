package main

import (
	"fmt"

	"github.com/Rohanpudasaini/learn-go/utils"
)

// var c, python, java bool
// var c, python, java = true, 12, "no!"

func main() {
	// var a, b int = 20, 0
	fmt.Println("Pointers in Go")
	val := 50
	fmt.Println("Before:", val, " at address: ", &val)
	utils.ChangeValue(&val)
	fmt.Println("After:", val, " at address: ", &val)
	// i := 12
	// f := float64(i)
	// u := uint(f)
	// utils.Something()
	// fmt.Println(c, python, java, i)
	// fmt.Printf("Type %T value: %v \n", f, f)
	// fmt.Println(u)
	// utils := utils.Calculator{A: 20, B: 10}
	// fmt.Println("Add: ", utils.Add())
	// fmt.Println("Subtract: ", utils.Subtract())
	// fmt.Println("Multiply: ", utils.Multiply())
	// var division, error = utils.Divide()
	// if error != nil {
	// 	fmt.Println("Error: ", error)
	// } else {
	// 	fmt.Println("Divide: ", division)
	// }
}
