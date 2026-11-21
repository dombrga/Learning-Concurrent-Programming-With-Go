package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/**
	8.1.4. Timing out on channels
**/

func sendMsgAfter(seconds time.Duration) <-chan string {
	messages := make(chan string)
	/** Comment:
		In t=2, when this goroutine wakes up, it will block on send
		since there is nothing to receive it anymore.
		In this this casem it does not matter since main exits anyway.
		Can be solved by a buffered channel of 1.
	**/
	go func() {
		time.Sleep(seconds)
		messages <- "Hello"
	}()
	return messages
}

func main() {
	t, _ := strconv.Atoi(os.Args[1])

	// This sends a message to the returned channel in 3 seconds.
	messages := sendMsgAfter(3 * time.Second)
	timeoutDuration := time.Duration(t) * time.Second

	fmt.Printf("Waiting for message for %d seconds...\n", t)

	select {
	case msg := <-messages:
		fmt.Println("Message received:", msg)
	case tNow := <-time.After(timeoutDuration):
		fmt.Println("Timed out. Waited until:", tNow.Format("15:04:05"))
	}
}
