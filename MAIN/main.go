package main

import (
	"fmt"

	colors "github.com/spencerguo11/go_learning_spencer/COLOR"
	functions "github.com/spencerguo11/go_learning_spencer/FUNCTION"
	structs "github.com/spencerguo11/go_learning_spencer/STRUCT"
)

func main() {
	fmt.Println(colors.ColorRed, "Hello world!")
	fmt.Println(colors.ColorReset) // Reset the color after printing
	// Add numbers
	AddRes := functions.Add(4, 8)
	fmt.Println("4 plus 8 is equal to ", AddRes, "!!")

	// Divide
	res1, err := functions.Divide(9, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println("9 divides by 3 is equal to ", res1)
	/*
		res2, err := functions.Divide(10, 0)
		if err != nil {
			panic(err) // panic function will stop and print the err message.
		}
		fmt.Println("9 divides by 3 is equal to ", res2)
	*/

	// Create Data Point
	point := structs.Point{X: 5, Y: 10}
	fmt.Println("Our Point is clocate at X=", point.X, "; Y=", point.Y)

	// Create Rectange + Compute Size
	rectangle := structs.Rectangle{Height: 14.12, Width: 19.05}
	size := rectangle.GetSize()
	fmt.Println("The rectangle of the height", rectangle.Height, "and width", rectangle.Width, "is", size, "in size\n")

}
