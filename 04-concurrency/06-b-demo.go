/* Using channel for communicating the result */
package main

import (
	"fmt"
)

func main() {

	// share memory by communicating
	ch := make(chan int)

	go add(100, 20, ch)
	result := <-ch //receiving result from the channel
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result //sending result to the channel
}
