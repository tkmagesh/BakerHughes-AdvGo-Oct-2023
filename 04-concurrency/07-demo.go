// channel behaviors

package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100

}
