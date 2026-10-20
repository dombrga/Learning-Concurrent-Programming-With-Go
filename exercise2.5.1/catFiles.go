package main

import (
	"fmt"
	"os"
	"time"
)

func sayHello() {
	fmt.Println("hello")
}

// func main() {
// 	fmt.Println("Number of CPUs:", runtime.NumCPU())
// 	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

// 	go sayHello()
// 	runtime.Gosched()
// 	fmt.Println("finished")
// }

// 2.5.1
func main() {
	args := os.Args
	fmt.Println("args: ", args)

	fileNames := args[1:]
	fmt.Println("fileNames: ", fileNames)

	for _, fileName := range fileNames {
		go func() {
			file, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println("err opening file:", err)
			}

			fmt.Println("cat file:", string(file))
		}()
	}

	time.Sleep(time.Microsecond * 1000)
}
