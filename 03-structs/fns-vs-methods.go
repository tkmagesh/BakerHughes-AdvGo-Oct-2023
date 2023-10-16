package main

import "fmt"

type Employee struct {
	Id int
	Name string
	Salary float32
} 
func main(){
	var i int = 100
	var iPtr *int
	iPtr = &i

	// how to access the value using the pointer
	fmt.Println(*iPtr /* dereferencing */)

	emp := Employee{
		Id : 100,
		Name : "Magesh",
		Salary : 10000,
	}
	fmt.Println(emp)

	var empPtr *Employee
	empPtr = &emp

	// Accessing the attributes using the pointer
	// fmt.Println((*empPtr).Id)
	fmt.Println(empPtr.Id)

	// As a function
	// AwardBonus(&emp, 2000)

	// As a method
	// (&emp).AwardBonusAsMethod(2000)
	// emp.AwardBonusAsMethod(2000)
	empPtr.AwardBonusAsMethod(2000)
	fmt.Println(emp.Salary)

}

func AwardBonus(emp *Employee, bonus float32){
	// update the emp salary with bonus
	// (*emp).Salary += bonus
	emp.Salary += bonus
}

func (emp *Employee) AwardBonusAsMethod( bonus float32){
	// update the emp salary with bonus
	// (*emp).Salary += bonus
	emp.Salary += bonus
}

