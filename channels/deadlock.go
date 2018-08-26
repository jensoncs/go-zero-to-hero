package main

import "fmt"

func main() {

	a := make(chan int)

	hello(a)
}

func hello(ch chan int) {
	fmt.Println(ch)
	ch <- 5
}
