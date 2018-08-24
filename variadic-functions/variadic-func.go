package main

import "fmt"

func find(num int, nums ...int) {

	found := false

	for _, v := range nums {
		if v == num {

			fmt.Println(num, nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, nums)
	}
}

func main() {
	nums := []int{89, 79, 99, 109}
	find(89, nums...)
	find(87)
	find(11, 22, 11)
}
