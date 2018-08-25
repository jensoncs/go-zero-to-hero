package main

import "fmt"

func main() {
	b := 255
	var a *int = &b
	fmt.Println("%f\n", a)
	fmt.Printf("%T\n", a)

	fmt.Println("*************")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*a)
	*a++
	fmt.Println(b)
}
