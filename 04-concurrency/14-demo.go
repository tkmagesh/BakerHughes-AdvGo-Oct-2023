package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 10
	}()

	go func() {
		time.Sleep(4 * time.Second)
		d3 := <-ch3
		fmt.Println("data from ch3 :", d3)
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 20
	}()

	for i := 0; i < 3; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case d2 := <-ch2:
			fmt.Println(d2)
		case ch3 <- 30:
			fmt.Println("data sent to ch3")
			/* default:
			fmt.Println("no channel operations were successful") */
		}
	}

}
