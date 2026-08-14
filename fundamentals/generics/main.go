package main

import "fmt"

/*

Key Takeaways:

- Generics let functions/types work with multiple data types.
- T is a type parameter
- [T any] means T can be any type
- Constraints restrict which types can be used.
- Generics reduce duplicate code while keeping type safety.
- Use generics when the same logic applies to other types.
- Dont' use generics when a normal function or interface is simpler.
*/


type Number interface {
	int | float64
}

func double[T Number](x T) T {
	return x * 2
}

func main() {
	fmt.Println("Welcome to generics")

	var intSlice = []int{1,2,3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1,2,3}
	fmt.Println(sumSlice[float32](float32Slice))

	printValue[int](42)
	printValue[string]("hello")

	pair[int, string](10, "hello there!")

	fmt.Println(first([]int{1,2,3}))
	fmt.Println(first([]string{"a","b","c"}))
}

func printValue[T any](value T) {
	fmt.Println(value)
}

func pair[T any, U any](a T, b U) {
	fmt.Println(a, b)
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}

	return sum
}

func first[T any](slice []T) T {
	return slice[0]
}

