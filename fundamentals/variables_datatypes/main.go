package main

import (
	"fmt"
)

func main() {
	
	
	// Default is 0
	var intNum int
	fmt.Println(intNum)

	// Specify 32 or 64 bit float type
	// Default is 0
	var floatNum float64
	fmt.Println(floatNum)

	// Adding a float and an integer together gives you a float back
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Printf("Type: %T\n", result)
	fmt.Println("The value is:", result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1 % intNum2)

	// Wah strings like abit weird ah

}