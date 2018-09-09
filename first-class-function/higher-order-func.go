package main

import "fmt"

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}
