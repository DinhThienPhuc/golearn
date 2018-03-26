package rectangle

import "math"

// Area function
func Area(length, width float64) float64 {
	return length * width
}

// Diagonal function
func Diagonal(length, width float64) float64 {
	return math.Sqrt(length*length + width*width)
}
