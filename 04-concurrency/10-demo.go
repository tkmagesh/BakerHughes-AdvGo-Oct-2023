package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {

	ch := generateNos()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(1 * time.Second)
			fmt.Println(data)
			continue
		}
		break
	}
	fmt.Println("Done")
}

// producer
func generateNos() <-chan int {
	ch := make(chan int)
	go func() {
		count := rand.Intn(20)
		fmt.Printf("Producer is about to produce %d values\n", count)
		for i := 1; i <= count; i++ {
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
