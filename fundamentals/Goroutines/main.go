package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Key Takeaways:

- A goroutine runs a function concurrently using the 'go' keyword.
- Use sync.WaitGroup to wait for goroutines to finish.
- Shared mutable data can cause race conditions when accessed concurrently
- Use a mutex to safely access shared mutable data
- Lock() protects writes; RLock() allows safe concurrent reads.
- Only lock the shared state that needs protection.
- Keep the critical section (code inside the lock) as small as possible
*/


var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}



func main() {
	fmt.Println("Welcome to GoRoutines")
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	// Stimulate db call delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}