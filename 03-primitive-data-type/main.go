package main

import "fmt"

func main(){
	// Integer Numbers
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)
	var conversion int32 = int32(integer8) + integer32
	fmt.Println(conversion)

	// Rune Data Type, store code that represent Unicode characters
	rune := '@'
	fmt.Println("Rune : ",rune)

	// Only positive numbers
	var integeruint uint = 20
	var integeruint8 uint8 =23
	fmt.Println(integeruint)
	fmt.Println(integeruint8)

	// Float Values
	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	fmt.Println(float32, float64)


	// Complex Numbers
	var complex128 complex128 = 10 + 8i
	fmt.Println("Complex Type: ", complex128)

}