package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4}
	fmt.Println(a)
	//	b := &a
	change(&a)
	fmt.Println(a)
}

func change(values *[4]int) {
	values[0] = (*values)[0]
	values[0] = 11
}
