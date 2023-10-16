/* Using channel for communicating the result */
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	// share memory by communicating
	ch := make(chan int)
	wg.Add(1)
	go add(100, 20, wg, ch)
	result := <-ch //receiving result from the channel
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result //sending result to the channel
	wg.Done()
}
