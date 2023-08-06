package main

import (
	"context"
	"fmt"
	"github.com/BoskyWSMFN/go-lesson-multithreading-architecture/internal/multithreading/simplechans"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	simplechans.SendNumbersCtx(ctx, wg, ch1)
	simplechans.PrintReceivedNumbers(wg, ch1)
	simplechans.SendNumbers(ch2)

	for i := range ch2 {
		fmt.Printf("main(): received number %d\n", i)
	}
}
