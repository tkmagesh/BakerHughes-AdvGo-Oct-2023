package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	// top level context
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)

	fmt.Println("Hit ENTER to shutdonw...")
	go func() {
		fmt.Scanln()
		cancel()
	}()

	wg.Add(1)
	go longRunningTask(cancelCtx, wg)

	wg.Wait()
}

func longRunningTask(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("Some work is being done")
			time.Sleep(500 * time.Millisecond)
		}

	}
}
