package barrier

import (
	"log"
	"sync"
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
	log.Println("called Wait")
	b.cond.L.Lock()
	b.waitCount++

	// If the # of expected goroutines are now all waiting, call Broadcast to continue execution.
	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast()
	} else {
		// Suspend the goroutine when # of waiting goroutines is less than barrier size.
		b.cond.Wait()
	}

	b.cond.L.Unlock()
}
