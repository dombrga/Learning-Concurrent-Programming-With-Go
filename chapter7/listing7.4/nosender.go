package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string)
	go sender(msgChannel)
	fmt.Println("Reading message from channel...")

	msg := <-msgChannel
	fmt.Println("Received:", msg)
}

// Deadlock error when a goroutine sends a message thru the channel
// and there is no receiver goroutine.
// Same deadlock error if there is a receiver but no sender.
func sender(messages chan string) {
	time.Sleep(5 * time.Second)
	fmt.Println("Sender slept for 5 secoonds")
}
