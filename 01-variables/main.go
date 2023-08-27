package main

import "fmt"

func main() {

	//Define Constants
	const pi float64 = 3.14

	// Define constants without date type
	const pi_two = 3.1415

	fmt.Println("PI ONE: ", pi)
	fmt.Println("PI TWO: ", pi_two)

	//Define integer variasbles
	base := 12
	var high int = 14
	var area int
	fmt.Println(base, high, area)

	// Zero values different data types and print default values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Example using variables
	const base_square = 10
	area_square := base_square * base_square
	fmt.Println("Area Square: ", area_square)

}
