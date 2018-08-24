package main

import "fmt"

func main() {

	a := make([]string, 4)
	fmt.Println(a)

	a[0] = "jenson"
	a[1] = "c"
	a[2] = "s"
	fmt.Println(a)
	fmt.Printf("length: %v, capacity: %v\n", len(a), cap(a))
	a = append(a, "gojek")
	fmt.Printf("length: %v, capacity: %v\n", len(a), cap(a))

	b := a[1:3]
	fmt.Println(b)
	fmt.Printf("length: %v, capacity: %v\n", len(b), cap(b))

	for _, v := range a {
		fmt.Println(v)
	}
}
