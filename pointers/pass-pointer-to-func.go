package main

import "fmt"

func main() {
	a := 58
	fmt.Println(a)
	b := &a
	chnage(b)
	fmt.Println(a)
}

func chnage(value *int) {
	*value = 55
}
