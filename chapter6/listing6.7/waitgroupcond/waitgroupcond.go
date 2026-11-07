package main

import (
	"fmt"
	"sync"
)

type WaitGroup struct {
	groupSize int
	cond      *sync.Cond
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{cond: sync.NewCond(&sync.Mutex{})}
}

func (wg *WaitGroup) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize = wg.groupSize + delta
	wg.cond.L.Unlock()
}

func (wg *WaitGroup) Wait() {
	wg.cond.L.Lock()

	// Suspend goroutines when there are outstanding/running goroutine.
	for wg.groupSize > 0 {
		wg.cond.Wait()
	}

	// wg.cond.Broadcast()
	wg.cond.L.Unlock()
}

func (wg *WaitGroup) Done() {
	wg.cond.L.Lock()
	wg.groupSize--

	// Broadcast when it is the last goroutine to resume suspended goroutines.
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}

	wg.cond.L.Unlock()
}

func main() {
	wg := NewWaitGroup()

	for i := 0; i < 5; i++ {
		go doWork(i, wg)
	}

	wg.Wait()
	fmt.Println("all completed")
}

func doWork(id int, wg *WaitGroup) {
	// i := rand.IntN(5)
	// time.Sleep(time.Duration(i) * time.Second)
	// fmt.Println(id, "Done working after", i, "seconds")
	wg.Done()
}
