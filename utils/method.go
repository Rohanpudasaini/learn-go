package utils

import errors "errors"

type Calculator struct {
	A int
	B int
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
