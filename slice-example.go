package main

import "fmt"

func main() {

	cars := []string{"ferrari", "honda", "ford"}

	fmt.Println(cars)
	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	cars = append(cars, "indica")
	fmt.Println(cars)
	fmt.Println(len(cars))
	fmt.Println(cap(cars))
	cars = append(cars, "a")
	fmt.Println(cars)
	fmt.Println(len(cars))
	fmt.Println(cap(cars))
	cars = append(cars, "b")
	fmt.Println(cars)
	fmt.Println(len(cars))
	fmt.Println(cap(cars))
	cars = append(cars, "c")
	fmt.Println(cars)
	fmt.Println(len(cars))
	fmt.Println(cap(cars))
}
