package main

import "fmt"

func main() {

	a := []int{8, 7, 6}
	fmt.Println(a)

	substract(a)
	fmt.Println(a)
}

func substract(number []int) {
	for i := range number {
		number[i] -= 2
	}
}
