package main

import "fmt"

type Employ struct {
	name, currency string
	salary         int
}

func main() {

	emp1 := Employ{
		name:     "ram",
		currency: "INR",
		salary:   100000,
	}
	emp1.display()
}

func (e Employ) display() {
	fmt.Println(e.name)
}
