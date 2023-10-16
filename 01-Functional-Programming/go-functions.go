package main

import "fmt"

func main(){
	/* 
	fn := func(){
		fmt.Println("fn invoked")
	} 
	*/
	var fn func()
	fn = func(){
		fmt.Println("fn invoked")
	} 
	fn()

	var add func(int, int)
	add = func(x,y int){
		fmt.Printf("Add Result : %d\n", x + y)
	}
	add(100,200)

	var multiply func(int, int) int
	multiply = func(x,y int) int {
		return x * y
	}
	fmt.Printf("Multiply Result : %d\n", multiply(100,200))

	greeter := getGreeter()
	greeter()
}

func getGreeter() func() {
	return func(){
		fmt.Println("Hello from greeter!")
	}
}