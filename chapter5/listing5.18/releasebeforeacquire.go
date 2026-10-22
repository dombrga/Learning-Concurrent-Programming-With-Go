package main

import (
	"fmt"

	listing5_16 "github.com/dombrga/learn-concurrency-go/chapter5/listing5.16"
)

// The code here is related and an improvement to listing5.6 using semaphores
// Two cases may happen each iteration:
//  1. Acquire runs first, which goes into the loop and parks at Wait().
//     Then doWork calls Release, and Acquire code continues on.
//  2. doWork runs first which ignores the for-loop-Wait() in Acquire and continues on.
func main() {
	semaphore := listing5_16.NewSemaphore(0)

	for i := 0; i < 50000; i++ {
		go doWork(semaphore)

		// fmt.Println("Waiting for child goroutine")
		semaphore.Acquire()
		// fmt.Println("Child goroutine")
	}
}

func doWork(semaphore *listing5_16.Semaphore) {
	fmt.Println("Work started")
	fmt.Println("Work finished")

	semaphore.Release()
}
