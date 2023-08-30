package main

// Import three modules
import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Define two variables
	var value_one int = 10
	var value_two int = 20

	// Using if
	if value_one == value_two {
		fmt.Println("It's equal")
	} else {
		fmt.Println("It is not equal")
	}

	// Use and
	if value_one == 10 && value_two == 20 {
		fmt.Println("The value one is equal to 10 and values two is equal to 20")
	}

	// Using or
	if value_one == 10 || value_two == 50 {
		fmt.Println("One value or two values is equal")
	}

	// Convert text to number and validate using if and else
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The number is: ", value)
	}

}
