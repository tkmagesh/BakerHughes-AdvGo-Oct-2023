package main

import "fmt"

type Product struct {
	Id int
	Name string
	Cost float32
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100-discountPercentage)/100)
}

// struct composition
type PerishableProduct struct {
	Product
	Id int
	Expiry string
}

// OVERRIDE the string method of the Product
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%v, Expiry = %q", pp.Product.String(), pp.Expiry)
}

func main(){
	// p := Product {10, "Pen", 5}
	
	p := Product {
		Id : 10,
		Name : "Pen",
		Cost : 5,
	}

	/* 
	var p2 *Product
	p2 = &p
	p2.Cost = 50
	fmt.Println(p)
	fmt.Println(p2)  
	*/
	
	fmt.Println(p)
	fmt.Println("After applying discount")
	p.ApplyDiscount(10)
	fmt.Println(p) 

	grapes := PerishableProduct {
		Product : Product {
			Id : 100,
			Name : "Grapes",
			Cost : 500,
		},
		Expiry : "2 Days",
	}

	fmt.Println(grapes)
	fmt.Println(grapes.Product.Id)
	fmt.Println(grapes.Id)

	/* grapes.ApplyDiscount(10)
	fmt.Println("After applying 10% discount")
	fmt.Println(grapes) */

	
}