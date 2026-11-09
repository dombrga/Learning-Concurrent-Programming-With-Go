package main

import (
	"fmt"
	"sync"
	"time"
)

func receiver(messages chan int, wg *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = <-messages
		fmt.Println("Received:", msg)
	}
	wg.Done()
}

func main() {
	msgChannel := make(chan int, 3)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go receiver(msgChannel, &wg)

	// The sender sends at a rate faster than the receiver receives the message.
	// There will be time when the sender is blocked because the buffer is full.
	// THe sender will be unblocked once the receiver receives a message.
	for i := 1; i <= 6; i++ {
		size := len(msgChannel)
		fmt.Printf("%s Sending: %d. Buffer size %d\n", time.Now().Format("15:04:05"), i, size)
		msgChannel <- i
		fmt.Println("after sending")
	}

	msgChannel <- -1
	wg.Wait()
}
