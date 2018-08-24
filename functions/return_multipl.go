package main

import (
	"fmt"
)

func returnmul(a, b int) (area, cir int) {
	area = a * b
	cir = a + b
	return
}

func main() {
	ar, cir := returnmul(3, 4)
	fmt.Println(ar)
	fmt.Println(cir)
}
