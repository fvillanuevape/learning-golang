package main

import "fmt"

func main() {

	// conditional for
	for index := 0; index <= 10; index++ {
		fmt.Println("Conditional for: ", index)
	}
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}

	// for while
	counter := 0
	for counter <= 10 {
		fmt.Println("For While: ", counter)
		counter++
	}

	// for forever
	counter_forever := 0
	for {
		fmt.Println("For forever: ", counter_forever)
		counter_forever++
	}

}
