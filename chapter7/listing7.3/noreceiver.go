package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string)
	go receiver(msgChannel)
	fmt.Println("Sending HELLO...")
	msgChannel <- "HELLO"
	fmt.Println("Sending THERE...")
	msgChannel <- "THERE"
	fmt.Println("Sending STOP...")
	msgChannel <- "STOP"
}

// Deadlock error when a goroutine sends a message thru the channel
// and there is no receiver goroutine.
// Same deadlock error if there is a receiver but no sender.
func receiver(messages chan string) {
	time.Sleep(5 * time.Second)
	fmt.Println("Receiver slept for 5 secoonds")
}
