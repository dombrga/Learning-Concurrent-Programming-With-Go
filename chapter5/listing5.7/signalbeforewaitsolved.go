package main

import (
	"fmt"
	"sync"
)

/*
Note: this program has a bug for demonstration purposes
We demonstrate how to fix this problem in the next listing
*/
// This corrects listing5.6 code.
// The main goroutine stops/blocks at Wait() until another goroutine calls, in this case, Signal().
func doWork(cond *sync.Cond) {
	fmt.Println("Work started")
	fmt.Println("Work finished")

	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()

	for i := 0; i < 50000; i++ {
		go doWork(cond)

		fmt.Println("Waiting for child goroutine")
		cond.Wait()
		fmt.Println("Child goroutine finished")
	}

	cond.L.Unlock()
}
