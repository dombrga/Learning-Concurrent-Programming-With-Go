package main

import (
	"fmt"

	listing5_16 "github.com/dombrga/learn-concurrency-go/chapter5/listing5.16"
)

type WaitGroup struct {
	sema *listing5_16.Semaphore
}

func NewWaitGroup(size int) *WaitGroup {
	return &WaitGroup{sema: listing5_16.NewSemaphore(1 - size)}
}

func (wg *WaitGroup) Wait() {
	wg.sema.Acquire()
}

func (wg *WaitGroup) Done() {
	wg.sema.Release()
}

// func (wg *WaitGroup) Add() {
// 	wg.sema.Acquire()
// }

func main() {
	wg := NewWaitGroup(5)

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
