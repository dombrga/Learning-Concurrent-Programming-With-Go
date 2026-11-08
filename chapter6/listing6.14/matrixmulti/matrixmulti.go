package main

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/dombrga/learn-concurrency-go/chapter6/listing6.10/barrier"
)

const matrixSize = 3

// From listing 6.14
func matrixMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; row++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[col][i]
			}
			result[row][col] = sum
		}
	}
}

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.IntN(10) - 5
		}
	}
}

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int, row int, barrier *barrier.Barrier) {
	log.Println("multi", row)
	for {
		// fmt.Println("multi 1")
		barrier.Wait()
		// fmt.Println("multi 2")
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}

			result[row][col] = sum
		}
		barrier.Wait()
	}
}

func main() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	barrier := barrier.NewBarrier(matrixSize + 1)

	for row := 0; row < matrixSize; row++ {
		go rowMultiply(&matrixA, &matrixB, &result, row, barrier)
	}

	for i := 0; i < 2; i++ {
		log.Println("generating")
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)
		// Needs to call this Wait to release all waiting goroutines.
		barrier.Wait()
		log.Println("before 2nd wait")
		// It is now turn for main goroutine to Wait until the row goroutines finish their multiplication and calls their Wait.
		barrier.Wait()

		// When the 3 row goroutines call Wait() plus the above Wait() of the main goroutine,
		// execution will continue here.
		for i := 0; i < matrixSize; i++ {
			fmt.Println(matrixA[i], matrixB[i], result[i])
		}
		fmt.Println()
		// fmt.Println(matrixA, matrixB, result)
	}
}
