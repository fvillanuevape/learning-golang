package main

import "fmt"

func main(){
	// Using Defer
	defer fmt.Println("It is Defer")
	fmt.Println("Firts execution")

	//Using Continue and break
	for index := 0; index < 10; index++ {
		fmt.Println(index)

		// Continue
		if index==2 {
			fmt.Println("Continue")
			continue
		}

		//Break
		if index==8 {
			fmt.Println("Break")
			break
		}
		
		
	}
}