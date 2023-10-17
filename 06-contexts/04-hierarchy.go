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

	overriddenRootKeyCtx := context.WithValue(ctx, "root-key", "new-root-value")
	oddValCtx := context.WithValue(overriddenRootKeyCtx, "odd-key", "odd-value")
	printOddTimeoutCtx, cancel := context.WithTimeout(oddValCtx, 3*time.Second)
	defer cancel()

	wg.Add(1)
	go printOdds(printOddTimeoutCtx, wg)

	// printEvenTimeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	printEvenTimeoutCtx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	wg.Add(1)
	go printEvens(printEvenTimeoutCtx, wg)
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

func printOdds(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[printOdds] root-key :", ctx.Value("root-key"))
	fmt.Println("[printOdds] odd-key :", ctx.Value("odd-key"))
LOOP:
	for no := 1; ; no += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("printOdds done!")
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("Odd :", no)
		}

	}
}

func printEvens(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := 0; ; no += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("printEvens done!")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Even :", no)
		}

	}
}
