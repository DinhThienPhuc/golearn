//geometry.go
package main

import (
	"fmt"
	"rectangle"
)

func main() {
	var length, width float64 = 3, 4
	fmt.Println("Geometrical shape properties")
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(length, width))
	fmt.Printf("Diagonal of the rectangle %.2f\n", rectangle.Diagonal(length, width))
}
