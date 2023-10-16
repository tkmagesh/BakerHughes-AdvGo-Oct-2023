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

// implement the PrintPerimeter function which can be used to print the perimiter of circle & rectangle

type PerimiterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimiterFinder){
	fmt.Println("Perimiter :", x.Perimeter())
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Breadth)
}


// 
/* 
func PrintShape(x interface {
	interface {
		Area() float32
	}
	interface {
		Perimeter() float32
	}
}){
	PrintArea(x)
	PrintPerimeter(x)
} 
*/

/* 
func PrintShape(x interface {
	AreaFinder
	PerimiterFinder
}){
	PrintArea(x)
	PrintPerimeter(x)
}  
*/

type ShapeStatsFinder interface {
	AreaFinder
	PerimiterFinder
}

func PrintShape(x ShapeStatsFinder){
	PrintArea(x)
	PrintPerimeter(x)
} 

// 
func (c Circle) String() string {
	return fmt.Sprintf("Circle : Radius : %v, Area : %v, Perimiter : %v", c.Radius, c.Area(), c.Perimeter())
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle : Length : %v, Breadth : %v, Area : %v, Perimiter : %v", r.Length, r.Breadth, r.Area(), r.Perimeter())
}


func main(){
	c := Circle{Radius : 12}
	// fmt.Println("Area :", c.Area())
	/* 
	PrintArea(c)
	PrintPerimeter(c) 
	*/

	// PrintShape(c)
	fmt.Println(c)

	r := Rectangle{Length: 10, Breadth:12}
	// fmt.Println("Area :", r.Area())
	/* 
	PrintArea(r)
	PrintPerimeter(r) 
	*/
	// PrintShape(r)
	fmt.Println(r)

	// PrintArea(100)
}

