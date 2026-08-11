package main

import (
	"fmt"
)

/*

Key Takeaways:

- Structs group related data together
- Methods are functions that belong to a specific type
- The receiver tells Go which type the method belongs to.

- func printUser(user User)       // function
- func (user User) printUser()    // method

- Go does not have traditional class inheritance
- Interfaces define behavior through method signatures
- A type implicitly implements an interface by having all of its methods
- Go favours composition over inheritance
- Struct embedding can promote fields and methods from the embedded type.

*/


type gasEngine struct {
	mpg uint8
	gallons uint8
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {

	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first")
	}

}



func main() {

	fmt.Println("fmt")

	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{"Zade"}}

	myEngine.mpg = 20
	fmt.Println(myEngine.gallons, myEngine.mpg, myEngine.ownerInfo.name)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())
	
	// Anonymous struct
	// Not reusable
	// var myEngine2 = struct{
	// 	mpg uint8
	// 	gallons uint8
	// }{25,20}
	
	fmt.Println()
}