package main

import "fmt"

type Employee struct {
	firstname, lastname string
	age, salary         int
}

func main() {

	emp1 := Employee{
		firstname: "sam",
		lastname:  "c",
		age:       22,
		salary:    100000,
	}
	emp2 := Employee{"george", "tom", 55, 800000}
	fmt.Println(emp1)
	fmt.Println(emp2)
	emp3 := Employee{}
	fmt.Println(emp3)
	emp3.firstname = "Jenson"
	fmt.Println(emp3)
	fmt.Println(emp1.firstname)
}
