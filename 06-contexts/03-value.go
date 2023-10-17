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
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")
	cancelCtx, cancel := context.WithCancel(valCtx)

	fmt.Println("Hit ENTER to stop")
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
	fmt.Println("root-key :", ctx.Value("root-key"))
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancellation signal received")
			break LOOP
		default:
			fmt.Println("Some work is being done")
			time.Sleep(500 * time.Millisecond)
		}

	}
}
