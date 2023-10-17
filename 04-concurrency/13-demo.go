package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 10
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 20
	}()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)

}
