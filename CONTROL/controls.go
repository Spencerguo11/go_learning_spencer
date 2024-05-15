package main

import (
	"fmt"
	"math"
)

func main() {
	year := 2023

	if year == 2023 {
		fmt.Println("The year is 2023!")
	} else if year == 2024 {
		fmt.Println("Is it already 2024!")
	} else {
		fmt.Println("The year is 2024!")
	}

	// Sell Ice Cream
	money := 1.2
	friend := true

	if money <= 2.0 || friend {
		fmt.Println("Here, have an ice cream cone!")
	} else {
		fmt.Println("Sorry, no money, no ice cream!")
	}

	// Square Numbers
	number := 8.0
	if squared := math.Pow(number, 2); squared > 16 {
		fmt.Println("Square of Number is greater than 16:", squared)
	}

	// For loop
	for i := 10; i >= 0; i-- {
		if i > 0 {
			fmt.Println("COUNTDOWN:", i)
		} else {
			fmt.Println("LEFTOFF!")
		}
	}

	// Range
	names := [5]string{"Peter", "Susan", "Joe", "Barbara", "Elizabeth"}

	for i, name := range names {
		fmt.Printf("Name #%d: %s\n", i, name)
	}

	// Switch
	favoriteColor := "purple"
	switch favoriteColor {
	case "red":
		fmt.Println("You like a very hot color")
	case "blue":
		fmt.Println("You like a very cool color")
	case "green":
		fmt.Println("Yuumy, Fruity!")
	default:
		fmt.Println("You like neither red,blue, or green???")
	}

}
