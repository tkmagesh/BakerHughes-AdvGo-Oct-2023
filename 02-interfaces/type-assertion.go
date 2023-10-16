package main

import "fmt"

type Product struct {
	Id int
	Name string
	Cost float32
}

func main(){
	// var x interface{}
	var x any
	x = 100
	x = true
	x = 99.99
	x = 10+5i
	x = "Pariatur ullamco ut culpa labore cillum labore nostrud occaecat dolore commodo magna cillum cillum aliquip."
	x = struct{}{}
	fmt.Println(x)

	// x = 100
	// x = "Reprehenderit laboris laborum cupidatat occaecat eiusmod."
	x = getExternalValue()
	// y := x.(int) + 200

	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("Non numeric data.. cannot be used for addition")
	}

	// x = "Enim excepteur sit eu nulla culpa ullamco et nisi consectetur officia."
	// x = 19+12i
	// x = true
	x = Product{100, "pen", 10}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 = ", val * 2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case complex128:
		fmt.Printf("x is a a complex128, real = %v & imag = %v\n", real(val), imag(val))
	case Product:
		fmt.Println("x is a product, x =", x);
	default :
		fmt.Println("x is an unknown type")
	}
	
}

func getExternalValue() interface{} {
	// data from external source
	return 100
	// return "Non elit do irure esse ad ad commodo proident ipsum tempor magna pariatur."
}