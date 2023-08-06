package simplechans

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func SendNumbers(ch chan int) {
	go func() {
		defer close(ch)

		for i := 1; i <= 5; i++ {
			ch <- i

			fmt.Printf("simplechans.SendNumbers(): sent number %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

const wgDelta = 1

func SendNumbersCtx(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	wg.Add(wgDelta)

	go func() {
		defer wg.Done()
		defer close(ch)

		for i := 0; true; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("simplechans.SendNumbersCtx(): context canceled\n")

				return
			case ch <- i:
				fmt.Printf("simplechans.SendNumbersCtx(): sent number %d\n", i)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func PrintReceivedNumbers(wg *sync.WaitGroup, ch chan int) {
	wg.Add(wgDelta)

	go func() {
		defer wg.Done()

		for i := range ch {
			fmt.Printf("simplechans.PrintReceivedNumbers(): received number %d\n", i)
		}

		fmt.Printf("simplechans.PrintReceivedNumbers(): channel closed\n")
	}()
}
