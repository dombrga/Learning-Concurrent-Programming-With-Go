package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	searchDir := args[len(args)-1]
	fmt.Println("searchDir: ", searchDir)

	commonFilePath, err := filepath.Abs(searchDir)
	if err != nil {
		panic("error absolute path")
	}
	fmt.Println("common file path:", commonFilePath)

	dirEntries, err := os.ReadDir(searchDir)
	if err != nil {
		panic("panic: error reading directory")
	}

	fmt.Println("dir entries:", dirEntries)

	searchStr := args[1]

	for _, file := range dirEntries {
		go func() {
			file, err := os.ReadFile(commonFilePath + "/" + file.Name())
			if err != nil {
				fmt.Println("err opening file:", err)
			}

			content := string(file)
			fmt.Println("cat file:", content)

			if strings.Contains(content, searchStr) {
				fmt.Println("contains a match")
			}
		}()
	}

	time.Sleep(time.Microsecond * 3000)
}
