package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	size      int // number of goroutines that will use the barrier
	waitCount int // number of waiting goroutines at the barrier
	cond      *sync.Cond
}

func NewBarrier(size int) *Barrier {
	return &Barrier{
		size:      size,
		waitCount: 0,
		cond:      sync.NewCond(&sync.Mutex{}),
	}
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount++

	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast()
	} else {
		// Suspend the goroutine when # of waiting goroutine is less than barrierSize.
		b.cond.Wait()
	}

	b.cond.L.Unlock()
}

func main() {
	barrier := NewBarrier(2)
	go workAndWait("Red", 2, barrier)
	go workAndWait("Blue", 5, barrier)
	time.Sleep(20 * time.Second)
}

func workAndWait(name string, timeToWork int, barrier *Barrier) {
	start := time.Now()
	for {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		barrier.Wait()
	}
}
