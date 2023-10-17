package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFib()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib() <-chan int {
	ch := make(chan int)
	timeOutCh := timeOut(7 * time.Second)

	go func() {
	LOOP:
		for x, y := 0, 1; ; {
			select {
			case <-timeOutCh:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

func timeOut(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		ch <- time.Now()
	}()
	return ch
}
