package main

import (
	"fmt"
	"math"

	functions "github.com/spencerguo11/go_learning_spencer/FUNCTION"
)

func main() {

	// Extract from Array
	numbers := [10]int{5, 17, 19, 32, 55, 14, 7, 21, 42, 90}
	val := numbers[1]
	fmt.Printf("The number on position 1 is: %d\n", val)
	section := numbers[3:6]
	fmt.Printf("A slice taken from the array: %v\n", section)

	// Extract from Map
	ages := map[string]float64{
		"John":    55.0,
		"Susan":   32.7,
		"Barbara": 23.2,
	}
	ageSusan := ages["Susan"]
	fmt.Printf("Susan is %.1f years old\n", ageSusan)

	ageUnknown := ages["Unknown"]
	fmt.Printf("This should not yield a useful result: %v\n", ageUnknown)

	/*
		ageUnknown, ok := ages["Unknown"]
		if !ok {
			panic("Map Key does not exist!")

		}
		fmt.Println("This shouldn't even print!")

	*/

	// Appending
	var container []float64

	for i := 2.0; i <= 100.0; i += math.Pow(i, 2) {
		container = append(container, i)
	}

	fmt.Println(container)
	fmt.Printf("Elements: %d; Space: %d\n", len(container), cap(container))

	// Loop Over Map
	for k, v := range ages {
		fmt.Printf("Name: %s, Age: %.1f\n", k, v)
	}

	var ageSum float64
	var elements float64

	for _, age := range ages {
		ageSum += age
		elements++
	}

	avgAge := ageSum / elements

	fmt.Printf("Average Age: %.1f\n", avgAge)

	// Pointers
	numbersValue := []float64{32.0, 55.0, 85.0}
	numbersPointer := &numbersValue

	fmt.Printf("Before: %v\n", numbersValue)
	functions.DivideByTwo(numbersPointer)
	fmt.Printf("After: %v\n", numbersValue)

}
