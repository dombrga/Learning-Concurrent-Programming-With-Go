package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	listing4_5 "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter4/listing4.5"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	mutex := sync.Mutex{}
	var frequency = make([]int, 26)

	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go func() {
			listing4_5.CountLetters(url, frequency, &mutex)
			wg.Done()
		}()
	}

	wg.Wait()
	mutex.Lock()

	for i, c := range listing4_5.AllLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	mutex.Unlock()
}

func doWork(id int, wg *sync.WaitGroup) {
	i := rand.IntN(5)
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println(id, "Done working after", i, "seconds")
	wg.Done()
}
