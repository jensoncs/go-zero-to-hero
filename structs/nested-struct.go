package main

import "fmt"

type Person struct {
	name    string
	age     int
	address Address
}

type Address struct {
	city, state string
}

func main() {

	var p Person
	p.name = "sam"
	p.age = 29
	p.address =
		Address{
			city:  "Bangalore",
			state: "KA",
		}

	fmt.Println(p.name, p.age, p.address)
}
