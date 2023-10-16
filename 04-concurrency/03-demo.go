package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 20, wg)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup) {
	result = x + y
	wg.Done()
}
