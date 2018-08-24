package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(value int) (diveide_by_two_output int, type_of_value bool) {
	diveide_by_two_output = value / 2
	type_of_value = value%2 == 0
	return
}
