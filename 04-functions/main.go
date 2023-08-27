package main

import "fmt"

// Define function
func firsFunction() {
	fmt.Println("My First Function")
}

//Define function with only arguments
func secondFunction(message string) {
	fmt.Println("Second Function with parameter: ", message)
}

// Define function with multiple arguments
func multipleArguments(a int, b int, c string) {
	fmt.Println("Function with multiple argument: ", a, b, c)
}

// Define function and return value
func returnValue(a int) int {
	return a * 2
}

//Define function return multiple values
func returnMultipleValues(a int, b string) (valueOne int, valueTwo string) {
	return a, b
}

// Calculate rectangle area
func calculateRectanguleArea(length, width int) int {
	rectangle_area := length * width
	return rectangle_area
}

func main() {

	firsFunction()

	secondFunction("My message")

	multipleArguments(10, 20, "My message")

	get_value := returnValue(100)
	fmt.Println(get_value)

	value_one, value_two := returnMultipleValues(100, "Value Two")
	fmt.Println("Value One: ", value_one)
	fmt.Println("Value Two: ", value_two)

	//Use "_" for use only value when function return multiple values

	val_one, _ := returnMultipleValues(500, " Val Two")
	fmt.Println(val_one)

	rectangle_area := calculateRectanguleArea(1000, 30)
	fmt.Println("Rectangle Area: ", rectangle_area)

}
