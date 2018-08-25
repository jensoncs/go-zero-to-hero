package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4}
	fmt.Println(a)
	change(a[:])
	fmt.Println(a)
}

func change(values []int) {
	values[0] = 12
}
