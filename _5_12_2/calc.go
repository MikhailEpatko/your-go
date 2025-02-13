package _5_12_2

import "errors"

func Calculate(a, b float64, op string) (res float64, err error) {
	switch op {
	case "add":
		res = a + b
	case "subtract":
		res = a - b
	case "multiply":
		res = a * b
	case "divide":
		if b == 0 {
			err = errors.New("division by zero")
		} else {
			res = a / b
		}
	default:
		err = errors.New("unknown operation")
	}
	return
}
