package simpleprint

import (
	"fmt"
	"sync"
	"time"
)

const wgDelta = 1

func PrintNumbers(wg *sync.WaitGroup) {
	wg.Add(wgDelta)

	go func() {
		defer wg.Done()

		for i := 1; i <= 5; i++ {
			fmt.Println("Number:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func PrintLetters(wg *sync.WaitGroup) {
	wg.Add(wgDelta)

	go func() {
		defer wg.Done()

		letters := []string{"A", "B", "C", "D", "E"}

		for _, letter := range letters {
			fmt.Println("Letter:", letter)
			time.Sleep(150 * time.Millisecond)
		}
	}()
}
