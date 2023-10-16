package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go f1() // handover this function to the scheduler to be scheduled for execution IN FUTURE
	}
	f2()
	// DO NOT DO THIS IN PRODUCTION
	// block the execution of the current function so that the scheduler can look for other goroutines scheduled and execute them
	// time.Sleep(6 * time.Second)

	// PRODUCTION Friendly
	wg.Wait() // block until the wg counter becomes 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
