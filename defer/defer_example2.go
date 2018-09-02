package main

import "fmt"

func print(a int) {
	fmt.Println("Value of a in defer function is ", a)
}

func main() {
	a := 5
	defer print(a)
	a = 10
	fmt.Println("Value of a defore defers is ", a)
}
