package utils

import (
	"errors"
	"fmt"
	"math"
)

var rectangle = Rectangle{12, 4}
var square = Square{12}
var circle = Circle{8}
var random = Random{}

func Init() {
	calculate(rectangle)
	calculate(square)
	calculate(circle)
	calculate(random)

}

func calculate(s Shape) {
	v2 := "Perimeter"
	fmt.Println()
	fmt.Print("Shape ")
	switch v := s.(type) {
	case Circle:
		fmt.Println("Circle")
		v2 = "Circumference"
	case Square:
		fmt.Println("Square")
	case Rectangle:
		fmt.Println("Rectangle")
	default:
		fmt.Printf("Invalid shape of type %T \n", v)
	}
	fmt.Println("Area: ", s.Area())
	fmt.Println(v2, " ", s.Perimeter())
	fmt.Println()
}

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

func (r Rectangle) Area() float64 {
	return r.Length * r.Bredth
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Bredth)
}

func (c Random) Area() float64 {
	return 1.02
}

func (r Random) Perimeter() float64 {
	return 2 * 1.003
}

func (s Square) Perimeter() float64 {
	return 4 * s.Length
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Calculator) Add() int {
	return c.A + c.B
}

func (c Calculator) Subtract() int {
	return c.A - c.B
}

func (c Calculator) Multiply() int {
	return c.A * c.B
}

func (c Calculator) Divide() (int, error) {
	if c.B == 0 {
		return 0, errors.New("Divide by zero Error") // handle division by zero
	}
	return c.A / c.B, nil
}
