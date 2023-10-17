package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	count := 10
	go generateNos(count, ch)
	for i := 1; i <= count; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func generateNos(count int, ch chan<- int) {
	for i := 1; i <= count; i++ {
		ch <- i * 10
		time.Sleep(1 * time.Second)
	}
}
