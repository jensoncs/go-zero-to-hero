package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5}
	b := a[1:4]
	fmt.Println(len(b))
	fmt.Println(cap(b))
	b = b[:cap(b)]
	fmt.Println(len(b))
	fmt.Println(cap(b))
}
