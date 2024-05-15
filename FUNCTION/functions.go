package functions

import "errors"

func Add(x, y int) int {
	result := x + y
	return result
}

func Divide(x, y float32) (float32, error) {
	// check fo Zero in Denominator
	if y == 0 {
		return 0.0, errors.New("Division by Zero")
	}

	// Calculate Division
	result := x / y
	// Return Result
	return result, nil
}
