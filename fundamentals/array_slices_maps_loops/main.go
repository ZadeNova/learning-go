package main

import (
	"fmt"
)

/*
Key Takeways:

Arrays:
- Fixed length and fixed type
- Indexable
- Stored Contiguously in memory

Slices:
- Dynamic, flexible view over an underlying array.
- Use append() to add elements.
- len() = number of elements.
- cap() = capacity of underlying array.
- make(type, length, capacity) creates a slice

Maps:
- Key-value data structure (hash table)
- Keys must be unique
- Access returns value + optional boolean indicating existence.
- delete() removes a key.

Loops:
- Go only has one loop: for
- range iterates over arrays, slices, and maps.
- for condition {} can be used like a while loop.
- Maps do not guarantee iteration order.
*/


func main(){
	fmt.Println("Arrays, Slices, Maps and Loops!")

	// Arrays
	// - Fixed Length
	// - Same Type Only
	// - Indexable
	// Contiguous In Memory

	//var intArr [3] int32
	// Another way to initialize the array
	//var intArr [3]int32 = [3]int32 {1,2,3}
	// OR
	intArr := [3]int32 {1,2,3}
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:3])

	// Print out the memory addresses. Difference of 4 bytes since int32 is 4 bytes.
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// Slices
	// Slices wrap arrays to give a more general, powerful, and convenient interfaces to sequences of data.
	fmt.Printf("START of SLICES \n\n")
	var intSlice []int32 = []int32{4,5,6}

	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Another way of making a slice is using make()
	// type, length, capacity
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3, cap(intSlice3))

	// Maps [ HashTable ]
	fmt.Printf("Start of Maps \n\n")
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	// Map returns two values. One optional value
	var age, ok = myMap2["BruceWayne"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name")
	}
	// Can delete keys as well
	// delete(myMap2, "Adam")

	// LOOPS
	fmt.Printf("START of LOOPS \n\n")
	// Maps in Golang are not sorted like python.
	for name, age := range myMap2 {

		fmt.Printf("Name: %v and Aged: %v \n", name, age)
	}
	
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// While loops
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i += 1
	}
	
	// Another way of doing the while loop
	// for {
	// 	if i >= 5 {
	// 		break
	// 	}

	// 	i += 1
	// }


	// the for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	



}
