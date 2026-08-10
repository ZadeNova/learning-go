package main

import (
	"errors"
	"fmt"
)

// Functions and Control structures
// Key Takeways:
// - Functions can return multiple values
// - Go uses error values instead of exceptions
// - Check errors with: if err != nil
// - if / else if / else for conditional logic
// - switch can replace complex if/else chains
// - switch can match a value or evaluate conditions
// - % returns the remainder. ( Modulo )
// - Functions can return named types (int,int, error)

func main() {

	var printValue string = "I love cats"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err  = intDivision(numerator, denominator)

	if err != nil {
		fmt.Printf(err.Error())

	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v\n", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v\n",result, remainder)
	}

	// Alternatively we can use switch statements

	switch {
		case err != nil:
			fmt.Printf(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v\n", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v\n",result, remainder)
	}

	// ANother way of using switch statements
	
	switch remainder {
	case 0:
		fmt.Printf("The division was exact\n")
	case 1,2:
		fmt.Printf("The division was close\n")
	default:
		fmt.Printf("The division was not close\n")
	}

	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int,int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero\n")

		return 0, 0, err
	}
	var result int = numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}