package main

import (
	"fmt"
)

// Variables and basic Data types

// Key Takeaways
// - Variables have zero values by default.
// - int, float32, float64, rune -> 0
// - string -> ""
// - bool -> false
// - GO is statically typed
// - := is the preferred local variable declaration
// - len(string) return bytes, not characters
// - rune is an alias for int32 (Unicode character)
// - Integer division truncates decimals
// - const defines immutable values.

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

	var myString string = `Hello World, I like cats`
	fmt.Println(myString)

	// YOu get the number of bytes
	fmt.Println(len("test"))

	var myRune rune = 'a'
	fmt.Println(myRune)


	var myBoolean bool = false
	fmt.Println(myBoolean)


	// Can initialize like that :=
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// Constants
	const myConst string = "const value"

	const pi float32 = 3.1415

}