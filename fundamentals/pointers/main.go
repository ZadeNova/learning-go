package main

import (
	"fmt"
)

/*
Key Takeaways:

- A pointer stores the memory address of a value
- & gets the address of a variable
- * accesses the value at an address
- Passing a value to a function creates a copy
- Passing a pointer allows the function to modify the original value
- Use pointers when you need to modify the original value or avoid copying large structs
- A pointer variable and the value it points to are stored separately
*/

func main() {

	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value if i is: %v\n", i)

	p = &i
	*p = 1

	fmt.Printf("\nThe value p points to is: %v\n", *p)
	fmt.Printf("The value if i is: %v\n", i)


	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("The result is: %v\n", result)
	fmt.Printf("The value of thing1 is: %v\n", thing1)


}

func square(thing2 *[5]float64) [5]float64 {

	fmt.Printf("The memory location of the thing2 array is: %p\n", thing2)
	for i := range thing2{
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}