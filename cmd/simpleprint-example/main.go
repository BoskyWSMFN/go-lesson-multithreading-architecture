package main

import (
	"fmt"
	"sync"

	"github.com/BoskyWSMFN/go-lesson-multithreading-architecture/internal/multithreading/simpleprint"
)

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Main: Starting goroutines...")

	simpleprint.PrintNumbers(wg)
	simpleprint.PrintLetters(wg)

	wg.Wait()

	fmt.Println("Main: Exiting...")
}
