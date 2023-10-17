/* Using channel for communicating the result */
package main

import (
	"fmt"
	"time"
)

// consumer
func main() {

	ch := add(100, 20)
	result := <-ch //receiving result from the channel
	fmt.Println(result)
}

// producer
func add(x, y int) <-chan int /* receive only channel */ {
	ch := make(chan int)
	go func() {
		time.Sleep(500 * time.Millisecond)
		result := x + y
		ch <- result //sending result to the channel
	}()
	return ch
}
