package main

import (
	"fmt"
	"time"
)

func main() {

	// Example switch case
	modulo := 10 % 2
	switch modulo {
	case 0:
		fmt.Println("It´s pair")
	default:
		fmt.Println("Is not pair")
	}

	// Example switch case initializer Statement
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("You must running")
	case today.Day() == 30:
		fmt.Println("Visit your family")
	case today.Day() == 31:
		fmt.Println(" You are swimming right now")
	}

	// without condition
	value:=200
	switch {
	case value >200:
		fmt.Println("It´s mayor")
	case value ==200:
		fmt.Println("It is equal")
	default:
		fmt.Println("Error.")
	}

}
