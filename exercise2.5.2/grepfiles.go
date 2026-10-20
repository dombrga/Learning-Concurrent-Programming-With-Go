package main

import (
	"fmt"
	"os"
	"strings"
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

	searchString := args[1]

	fileNames := args[2:]
	fmt.Println("fileNames: ", fileNames)

	for _, fileName := range fileNames {
		go func() {
			file, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println("err opening file:", err)
			}

			content := string(file)
			fmt.Println("cat file:", content)

			if strings.Contains(content, searchString) {
				fmt.Println("contains a match")
			}
		}()
	}

	time.Sleep(time.Microsecond * 1000)
}
