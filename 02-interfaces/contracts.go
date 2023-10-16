package main

import (
	"math"
	"fmt"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length float32
	Breadth float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Breadth
}

/* 
func PrintArea(x interface{}){
	if c, ok := x.(Circle); ok {
		fmt.Println("Area :", c.Area())
	}
	if r, ok := x.(Rectangle); ok {
		fmt.Println("Area :", r.Area())
	}
} 
*/

/* 
func PrintArea(x interface{}){
	if obj, ok := x.(interface{ Area() float32 }); ok {
		fmt.Println("Area :", obj.Area())
	} else {
		fmt.Println("Given object doesnot have an area method")
	}
} 
*/

/* 
func PrintArea(x interface{ Area() float32 }){
	fmt.Println("Area :", x.Area())	
} 
*/

type AreaFinder interface{ 
	Area() float32 
}

func PrintArea(x AreaFinder){
	fmt.Println("Area :", x.Area())	
}


func main(){
	c := Circle{Radius : 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)

	r := Rectangle{Length: 10, Breadth:12}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)

	PrintArea(100)
}

