package main

import "fmt"

func add(a, b int) int {
	sum := a + b
	return sum
}

func main() {
	fmt.Println(add(3, 4))
}
