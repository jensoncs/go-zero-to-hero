package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3}
	sum := 0

	for _, v := range array {
		fmt.Println(v)
		sum += v
	}
	fmt.Println(sum)
}
