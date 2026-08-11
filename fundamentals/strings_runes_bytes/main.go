package main

import (
	"fmt"
	"strings"
)

/*

Key Takeaways:

- Strings are immutable sequences of bytes.
- Indexing a string returns a byte
- Range over a string decodes UTF-8 and returns byte index + rune
- len(string) returns the number of bytes
- Use []rune when you need to work with unicode characters.
- Rune is an alias for Int32
- strings.Builder is more efficient for repeatedly concatenating strings.
myString[0]       // byte
range myString    // rune
[]rune(myString)  // runes
len(myString)     // bytes


*/
// Strings are an array of bytes

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("The length of 'myString' is %v\n", len(myString))

	// Runes
	fmt.Println("RUNES BEGIN HERE BELOW")
	var myString2 = []rune("résumé")
	var indexed2 = myString2[1]
	fmt.Printf("%v, %T\n", indexed2, indexed2)

	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	var strSlice = []string{"j","o","b","l","e","s","s"}
	var strBuilder strings.Builder
	var catStr = ""
	
	for i := range strSlice{
		// Creating a new string everytime.
		catStr += strSlice[i]
		strBuilder.WriteString(strSlice[i])
	}

	// Strings are immutable in Golang
	fmt.Println(strSlice)
	fmt.Printf("\n%v\n",catStr)
	fmt.Printf("\n%v\n", strBuilder.String())

}
