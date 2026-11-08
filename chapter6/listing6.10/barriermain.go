package main

import (
	"fmt"
	"time"

	"github.com/dombrga/learn-concurrency-go/chapter6/listing6.10/barrier"
)

func main() {
	barrier := barrier.NewBarrier(2)
	go workAndWait("Red", 2, barrier)
	go workAndWait("Blue", 5, barrier)
	time.Sleep(20 * time.Second)
}

func workAndWait(name string, timeToWork int, barrier *barrier.Barrier) {
	start := time.Now()
	for {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		barrier.Wait()
	}
}
