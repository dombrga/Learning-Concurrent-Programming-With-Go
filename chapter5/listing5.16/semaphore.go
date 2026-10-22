package listing5_16

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		permits: n,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (rw *Semaphore) Acquire() {
	fmt.Println("acquiring")
	rw.cond.L.Lock()

	for rw.permits <= 0 {
		rw.cond.Wait()
	}

	rw.permits--
	rw.cond.L.Unlock()
}

func (rw *Semaphore) Release() {
	rw.cond.L.Lock()

	rw.cond.Signal()
	rw.permits++

	rw.cond.L.Unlock()
}
