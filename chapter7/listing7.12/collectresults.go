package main

import "fmt"

func findFactors(number int) []int {
	fmt.Println("finding factors for", number)
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	resultCh := make(chan []int)

	go func() {
		resultCh <- findFactors(3419110721)
	}()

	fmt.Println("after goroutine")
	fmt.Println(findFactors(478))
	fmt.Println("after second")

	// This print will not run until findFactors(3419110721) is finished
	// because "receiving" is a blocking operation until a value is sent.
	fmt.Println(<-resultCh)
}
