package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fileSearch(dir string, filename string, wg *sync.WaitGroup, mu *sync.Mutex, cont *[]string) {
	files, _ := os.ReadDir(dir)

	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			mu.Lock()
			fmt.Println(fpath)
			*cont = append(*cont, file.Name())
			mu.Unlock()
		}

		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg, mu, cont)
		}
	}

	wg.Done()
}

func main() {
	mu := sync.Mutex{} // Exercise 6.3.1
	wg := sync.WaitGroup{}
	cont := []string{}

	wg.Add(1)
	go fileSearch(os.Args[1], os.Args[2], &wg, &mu, &cont)
	wg.Wait()

	// Do alphabetization here using cont...
	fmt.Println("matches:", cont)
}
