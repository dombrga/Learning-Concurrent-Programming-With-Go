package main

import (
	"fmt"
)

const (
	passwordToGuess = "go far"
	alphabet        = " abcdefghijklmnopqrstuvwxyz"
)

func toBase27(n int) string {
	result := ""
	for n > 0 {
		result = string(alphabet[n%27]) + result
		n /= 27
	}
	return result
}

func guessPassword(from int, upto int, stop chan int, result chan string) {
	for guessN := from; guessN < upto; guessN++ {
		select {
		case <-stop:
			fmt.Printf("Stopped at %d [%d, %d)\n", guessN, from, upto)
			return
		default:
			guess := toBase27(guessN)
			if guess == passwordToGuess {
				result <- guess
				close(stop)
				return
			}
		}
	}
}

func main() {
	finished := make(chan int) // common channel used to signal that password is found
	passwordFound := make(chan string)

	for i := 1; i <= 387_420_488; i += 10_000_000 {
		go guessPassword(i, i+10_000_000, finished, passwordFound)
	}

	fmt.Println("goroutine creations are done, waiting for password match")
	fmt.Println("password found:", <-passwordFound)
	close(passwordFound)
	// time.Sleep(5 * time.Second)
}
