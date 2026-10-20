package main

import (
	"fmt"
	"time"

	listing4_12 "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter4/listing4.12"
)

func main() {
	rwMutex := listing4_12.ReadWriteMutex{}

	for i := 0; i < 2; i++ {
		go func() {
			for {
				rwMutex.ReadLock()
				time.Sleep(time.Second * 1)

				fmt.Println("Read clone")
				rwMutex.ReadUnlock()
			}
		}()
	}

	time.Sleep(time.Second * 1)
	rwMutex.WriteLock()
	fmt.Println("Write finished")
}
