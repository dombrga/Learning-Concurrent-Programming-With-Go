package main

import (
	"container/list"
	"sync"

	listing5_16 "github.com/dombrga/learn-concurrency-go/chapter5/listing5.16"
)

type Channel[M any] struct {
	capacitySema *listing5_16.Semaphore // semaphore to block sender when the buffer is full
	sizeSema     *listing5_16.Semaphore // semaphore to block the receiver when buffer is empty
	mutex        sync.Mutex             // Mutex to protect our shared list data structure
	buffer       *list.List             // Linked list to be used as queue data structure
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		capacitySema: listing5_16.NewSemaphore(capacity),
		sizeSema:     listing5_16.NewSemaphore(0),
		buffer:       list.New(),
		// mutex:        sync.Mutex{},
	}
}

func (c *Channel[M]) Send(message M) {
	c.capacitySema.Acquire()

	c.mutex.Lock()
	c.buffer.PushBack(message)
	c.mutex.Unlock()

	c.sizeSema.Release()
}

func (c *Channel[M]) Receive() M {
	c.capacitySema.Release()

	c.sizeSema.Acquire()

	c.mutex.Lock()
	v := c.buffer.Remove(c.buffer.Front()).(M)
	c.mutex.Unlock()

	return v
}
