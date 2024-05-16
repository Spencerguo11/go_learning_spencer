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

// Divide by Two (Overwriting Elements in Slice)

func DivideByTwo(pointer *[]float64) {
	for i, val := range *pointer { //  using a pointer to extract the value
		(*pointer)[i] = val / 2.0
	}
}
