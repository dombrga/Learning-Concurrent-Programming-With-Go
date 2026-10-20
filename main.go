package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	fmt.Println("hello")
}

// func main() {
// 	fmt.Println("Number of CPUs:", runtime.NumCPU())
// 	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

// 	go sayHello()
	runtime.Gosched()
// 	fmt.Println("finished")
// }

// 2.5.1
func main() {

}
