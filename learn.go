package main

import "fmt"

func main() {
	x := 3
	y := 7
	fmt.Println(x, y)

	x = x + y
	fmt.Println(x, y)

	y = x - y
	fmt.Println(x, y)

	x = x - y
	fmt.Println(x, y)
}
