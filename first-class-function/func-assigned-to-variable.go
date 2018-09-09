package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("Hello world")
	}
	a()
	fmt.Printf("%T", a)
}
