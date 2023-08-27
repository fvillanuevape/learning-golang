package main

import "fmt"

func main() {

	x := 50
	y := 10

	//Addition
	result := x + y
	fmt.Println(result)

	// Subtraction
	result = x - y
	fmt.Println(result)

	// Multiplication
	result = x * y
	fmt.Println(result)

	//Division
	result = x / y
	fmt.Println(result)

	// Module
	result = x % y
	fmt.Println(result)

	// Increment and Decrement
	x++
	fmt.Println("x: ", x)
	x--
	fmt.Println("x: ", x)

	// Assignment Operators
	x += 10
	fmt.Println(x)
	x *= 10
	fmt.Println(x)

	// Example calculate rectangle area
	rectangle_length := 50
	rectangle_width := 30
	rectangle_area := rectangle_length * rectangle_width
	fmt.Println("Rectangle Area: ", rectangle_area)

}
