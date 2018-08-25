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
	emp := &emp1
	fmt.Println(&emp)
	fmt.Println(*emp)
}
