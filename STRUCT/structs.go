// Declare Package
package structs

// Define Point
type Point struct {
	X int
	Y int
}

// Define Rectangle
type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) GetSize() float32 {
	// Calculate Size
	size := r.Height * r.Width

	// Return Size
	return size
}
