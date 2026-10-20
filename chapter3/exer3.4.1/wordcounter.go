package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

/*
Note: this program has a race condition for demonstration purposes
Additionally we have a timer at the end which you might need to adjust
depending on how fast your internet connection is.
In later chapters we cover how to wait for threads to complete their work
*/
func countWords(url string, frequency map[string]int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	words := strings.Fields(string(body[0:]))
	fmt.Println("body:", len(words))

	for _, b := range words {
		w := strings.ToLower(string(b))
		freq, ok := frequency[w]
		if !ok {
			frequency[w] = 1
		}

		frequency[w] = freq + 1
		// cIndex := strings.Index(allLetters, c)
		// if cIndex >= 0 {
		// 	frequency[cIndex] += 1
		// }
	}
	fmt.Println("Completed:", url)
}

func main() {
	// var frequency = make([]int, 26)
	var wordFrequency = make(map[string]int)

	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, wordFrequency)
	}

	time.Sleep(3 * time.Second)

	// for i, c := range allLetters {
	// 	fmt.Printf("%c-%d ", c, frequency[i])
	// 	fmt.Println()
	// }
	fmt.Println(wordFrequency, len(wordFrequency))
}
