package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5}

	b := a[1:4]

	fmt.Println(a)

	for i, v := range b {
		i++
	}
	fmt.Println(a)
}
