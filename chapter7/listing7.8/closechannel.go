package main

import (
	"fmt"
	"time"
)

func receiver(messages <-chan int) {
	for {
		// the "more" flag tells whether the channel is closed.
		msg, more := <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received", msg)
		time.Sleep(1 * time.Second)
		if !more {
			return
		}
	}
}

func anotherReceiver(messages <-chan int) {
	// This syntax allows you to read all messages
	// from a channel until it is closed.
	for msg := range messages {
		fmt.Println(time.Now().Format("15:04:05"), "Received", msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	msgChannel := make(chan int)
	// go receiver(msgChannel)
	go anotherReceiver(msgChannel)

	for i := 1; i <= 3; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending", i)
		msgChannel <- i
		time.Sleep(1 * time.Second)
	}
	close(msgChannel)

	// The following code will result to error since the channel
	// is already closed and it is still sending message to
	// the channel.
	// for i := 1; i <= 3; i++ {
	// 	fmt.Println(time.Now().Format("15:04:05"), "Sending", i)
	// 	msgChannel <- i
	// 	time.Sleep(1 * time.Second)
	// }

	time.Sleep(3 * time.Second)
}
