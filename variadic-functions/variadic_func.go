package main

import "fmt"

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println(sum)
}
