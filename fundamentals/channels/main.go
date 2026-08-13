package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Channels:

- Channels allow goroutines to communicate with each other.
- <- [sends a value into a channel]
- <-channel [receives a value from a channel]
- Sending/receiving on an unbuffered channel blocks until the other side is ready.
- select waits for the first available channel operation.
*/


var MAX_CHICKEN_PRICE float32 = 5
var MAX_BEEF_PRICE float32 = 8

func main() {
	fmt.Println("Welcome to Go Channels")

	var chickenChannel = make(chan string)
	var beefChannel = make(chan string)
	var websites = []string{"ntuc.com", "shengshiong.com", "walmart.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkBeefPrices(websites[i], beefChannel)
	}
	
	sendMessage(chickenChannel, beefChannel)

}

func checkBeefPrices(website string, beefChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var beefPrice = rand.Float32() * 20
		if beefPrice <= MAX_BEEF_PRICE {
			beefChannel <- website
			break
		}
	}
}



func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, beefChannel chan string) {
	
	select {
	case website := <-chickenChannel:
		fmt.Printf("\n Text Sent: Found deal on chicken at %v", website)
	case website := <-beefChannel:
		fmt.Printf("\nEmail Sent: Found deal on beef at %v", website)
	}

}