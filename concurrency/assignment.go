package main

import "fmt"

func square(a, b, c int, sq chan int) {
	out := (a * a) + (b * b) + (c * c)
	sq <- out
}

func cube(a, b, c int, cu chan int) {
	out := (a * a * a) + (b * b * b) + (c * c * c)
	cu <- out
}

func main() {
	sq := make(chan int)
	cu := make(chan int)

	go square(1, 2, 3, sq)
	go cube(1, 2, 3, cu)
	x := <-sq
	y := <-cu

	output := x + y
	fmt.Println(output)
}
